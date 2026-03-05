package billingapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type PdfRequest struct {
	HtmlContent string `json:"htmlContent"`
}

func DownloadPdf(w http.ResponseWriter, r *http.Request) {
	log.Println("DownloadPdf (+)")
	defer log.Println("DownloadPdf (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parse Request Body
	var req PdfRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.HtmlContent == "" {
		http.Error(w, "htmlContent is required", http.StatusBadRequest)
		return
	}

	// 2. Create Temp HTML File
	tmpDir := os.TempDir()
	timestamp := time.Now().Format("20060102150405")
	htmlFilename := filepath.Join(tmpDir, fmt.Sprintf("bill_%s.html", timestamp))

	// Use os.WriteFile (modern replacement for ioutil.WriteFile)
	err = os.WriteFile(htmlFilename, []byte(req.HtmlContent), 0644)
	if err != nil {
		log.Println("Error writing temp HTML file:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer os.Remove(htmlFilename) // Cleanup HTML file

	// 3. Generate PDF using Chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Convert path to proper file URL for Windows
	absPath, err := filepath.Abs(htmlFilename)
	if err != nil {
		log.Println("Error getting absolute path:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert backslashes to forward slashes strictly for the URL
	// Note: Windows paths in file:// URLs should look like file:///C:/path/to/file
	fileUrl := "file:///" + filepath.ToSlash(absPath)
	log.Printf("Navigating to: %s\n", fileUrl)

	var pdfBuffer []byte

	// Create a timeout context to prevent hanging
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err = chromedp.Run(ctx,
		chromedp.Navigate(fileUrl),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			// Generate the PDF
			// WithPrintBackground(true) ensures colors/styles are preserved
			pdfBuffer, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.27).   // A4 Width in inches
				WithPaperHeight(11.69). // A4 Height in inches
				Do(ctx)
			return err
		}),
	)

	if err != nil {
		log.Printf("Error generating PDF with chromedp: %v", err)
		http.Error(w, fmt.Sprintf("PDF Generation Failed: %v", err), http.StatusInternalServerError)
		return
	}

	// 4. Serve PDF directly from memory (no need to write to disk and read back)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"invoice_%s.pdf\"", timestamp))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(pdfBuffer)))
	w.Write(pdfBuffer)
}

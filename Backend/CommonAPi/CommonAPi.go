package commonapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetBankDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBankDetails (+)")
	defer log.Println("GetBankDetails (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	ifsc := r.URL.Query().Get("ifsc")
	if ifsc == "" {
		http.Error(w, "IFSC code is required", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://ifsc.razorpay.com/%s", ifsc)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching bank details:", err)
		http.Error(w, "Error fetching bank details", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

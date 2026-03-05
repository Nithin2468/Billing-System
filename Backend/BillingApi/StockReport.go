package billingapi

import (
	"billingapp/gwdb"
	"encoding/json"
	"log"
	"net/http"
)

type StockReportItem struct {
	Name       string  `json:"name"`
	BatchID    string  `json:"batchId"`
	ExpiryDate string  `json:"expiryDate"`
	CurrentQty int     `json:"currentQty"`
	SoldQty    int     `json:"soldQty"`
	TotalSales float64 `json:"totalSales"`
}

func GetStockReport(w http.ResponseWriter, r *http.Request) {
	log.Println("GetStockReport (+) - Request Received")
	defer log.Println("GetStockReport (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	// Base query for medicines (Current Stock)
	// We need to LEFT JOIN with bill_items to get sales within date range
	// This is a bit complex because bill_items needs to be filtered by date.
	// Alternative: Fetch all medicines, then for each, query sales. Or use a subquery.

	// Better approach:
	// 1. Get all medicines with current stock.
	// 2. Get aggregated sales from bill_items within date range, grouped by medicine_name and batch_id.
	// 3. Merge in Go.

	// Step 1: Fetch Medicines
	medQuery := `SELECT name, batch_id, expiry_date, quantity FROM medicines`
	rows, err := gwdb.GDbconn.Query(medQuery)
	if err != nil {
		log.Println("Error fetching medicines:", err)
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	stockMap := make(map[string]*StockReportItem)

	for rows.Next() {
		var item StockReportItem
		if err := rows.Scan(&item.Name, &item.BatchID, &item.ExpiryDate, &item.CurrentQty); err != nil {
			continue
		}
		key := item.Name + "|" + item.BatchID
		stockMap[key] = &item
	}
	rows.Close()

	// Step 2: Fetch Sales if date range provided
	if startDate != "" && endDate != "" {
		log.Printf("Fetching sales from %s to %s\n", startDate, endDate)

		// Join with bills table to filter by bill_date (created_at)
		salesQuery := `
			SELECT bi.medicine_name, bi.batch_id, SUM(bi.quantity), SUM(bi.amount)
			FROM bill_items bi
			JOIN bills b ON bi.bill_number = b.bill_number
			WHERE DATE(b.created_at) BETWEEN ? AND ?
			GROUP BY bi.medicine_name, bi.batch_id
		`

		sRows, err := gwdb.GDbconn.Query(salesQuery, startDate, endDate)
		if err != nil {
			log.Println("Error fetching sales:", err)
			// Don't fail entire request, just log
		} else {
			defer sRows.Close()
			for sRows.Next() {
				var name, batch string
				var qty int
				var amount float64
				if err := sRows.Scan(&name, &batch, &qty, &amount); err != nil {
					continue
				}

				key := name + "|" + batch
				if item, exists := stockMap[key]; exists {
					item.SoldQty = qty
					item.TotalSales = amount
				} else {
					// Edge case: Medicine might be deleted from medicines table but exist in history?
					// Or maybe just show it as 0 current stock.
					stockMap[key] = &StockReportItem{
						Name: name, BatchID: batch, CurrentQty: 0, SoldQty: qty, TotalSales: amount,
					}
				}
			}
		}
	}

	// Convert map to slice
	var report []StockReportItem
	for _, item := range stockMap {
		report = append(report, *item)
	}

	log.Printf("Returning report with %d items\n", len(report))
	json.NewEncoder(w).Encode(report)
}

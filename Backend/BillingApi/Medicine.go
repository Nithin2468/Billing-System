package billingapi

import (
	"billingapp/gwdb"
	"encoding/json"
	"log"
	"net/http"
)

type Medicine struct {
	ID         int     `json:"id,omitempty"`
	Name       string  `json:"name"`
	BatchID    string  `json:"batchId"`
	ExpiryDate string  `json:"expiryDate"`
	Quantity   int     `json:"qty"`
	MRP        float64 `json:"mrp"`
	GST        float64 `json:"gst"`
}

type MedicineResponse struct {
	Message string     `json:"message"`
	Status  bool       `json:"status"`
	Data    []Medicine `json:"data,omitempty"`
}

// AddMedicine inserts a new medicine into the database
func AddMedicine(w http.ResponseWriter, r *http.Request) {
	log.Println("AddMedicine (+) - Request Received")
	defer log.Println("AddMedicine (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req Medicine
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Adding Medicine: %+v\n", req)

	query := `INSERT INTO medicines (name, batch_id, expiry_date, quantity, mrp, gst) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := gwdb.GDbconn.Exec(query, req.Name, req.BatchID, req.ExpiryDate, req.Quantity, req.MRP, req.GST)
	if err != nil {
		log.Println("DB Error inserting medicine:", err)
		http.Error(w, "Failed to add medicine", http.StatusInternalServerError)
		return
	}

	log.Println("Medicine added successfully")
	json.NewEncoder(w).Encode(MedicineResponse{
		Status:  true,
		Message: "Medicine added successfully",
	})
}

// GetMedicines fetches all medicines from the database
func GetMedicines(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMedicines (+) - Fetching Stock")
	defer log.Println("GetMedicines (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	query := `SELECT id, name, batch_id, expiry_date, quantity, mrp, gst FROM medicines`
	rows, err := gwdb.GDbconn.Query(query)
	if err != nil {
		log.Println("DB Error fetching medicines:", err)
		http.Error(w, "Failed to fetch medicines", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var medicines []Medicine
	for rows.Next() {
		var m Medicine
		if err := rows.Scan(&m.ID, &m.Name, &m.BatchID, &m.ExpiryDate, &m.Quantity, &m.MRP, &m.GST); err != nil {
			log.Println("Error scanning medicine row:", err)
			continue
		}
		medicines = append(medicines, m)
	}

	log.Printf("Fetched %d medicines\n", len(medicines))
	json.NewEncoder(w).Encode(medicines)
}

// UpdateMedicine updates stock quantity
func UpdateMedicine(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateMedicine (+) - Request Received")
	defer log.Println("UpdateMedicine (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req Medicine
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Updating Medicine: %+v\n", req)

	query := `UPDATE medicines SET quantity = ? WHERE name = ? AND batch_id = ?`
	_, err := gwdb.GDbconn.Exec(query, req.Quantity, req.Name, req.BatchID)
	if err != nil {
		log.Println("DB Error updating medicine:", err)
		http.Error(w, "Failed to update medicine", http.StatusInternalServerError)
		return
	}

	log.Println("Medicine updated successfully")
	json.NewEncoder(w).Encode(MedicineResponse{
		Status:  true,
		Message: "Medicine updated successfully",
	})
}

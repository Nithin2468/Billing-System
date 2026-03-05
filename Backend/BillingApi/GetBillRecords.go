package billingapi

import (
	"billingapp/gwdb"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type BillData struct {
	BillNumber   string        `json:"billNumber"`
	Recipient    Recipient     `json:"recipient"`
	Consignee    Consignee     `json:"consignee"`
	Invoice      Invoice       `json:"invoice"`
	Bank         Bank          `json:"bank"`
	BillingItems []BillingItem `json:"billingItems"`
	SubTotal     float64       `json:"subTotal"`
	Discount     Discount      `json:"discount"`
	GrandTotal   float64       `json:"grandTotal"`
	PartyName    string        `json:"partyName"`
	Amount       float64       `json:"amount"` // Legacy field? Keeping it for now.
}

type Recipient struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Gstin   string `json:"gstin"`
}

type Consignee struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Gstin     string `json:"gstin"`
	DrugLicNo string `json:"drugLicNo"`
}

type Invoice struct {
	GstInvoiceNo      string `json:"gstInvoiceNo"`
	GstInvoiceDate    string `json:"gstInvoiceDate"`
	ReverseCharge     string `json:"reverseCharge"`
	SupplierRefNo     string `json:"supplierRefNo"`
	PaymentTerms      string `json:"paymentTerms"`
	DeliveryTerms     string `json:"deliveryTerms"`
	ContactPerson     string `json:"contactPerson"`
	ContactNumber     string `json:"contactNumber"`
	Email             string `json:"email"`
	DispatchedThrough string `json:"dispatchedThrough"`
	Destination       string `json:"destination"`
	Remark            string `json:"remark"`
}

type Bank struct {
	AccountNumber string `json:"accountNumber"`
	IfscCode      string `json:"ifscCode"`
	BankName      string `json:"bankName"`
	Branch        string `json:"branch"`
}

type BillingItem struct {
	MedicineName string  `json:"medicineName"`
	BatchId      string  `json:"batchId"`
	ExpiryDate   string  `json:"expiryDate"`
	Quantity     float64 `json:"quantity"`
	Mrp          float64 `json:"mrp"`
	Gst          float64 `json:"gst"`
	Amount       float64 `json:"amount"`
}

type Discount struct {
	Type   string  `json:"type"`
	Value  float64 `json:"value"`
	Amount float64 `json:"amount"`
}

func BillDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("BillDetails (+)")
	defer log.Println("BillDetails (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var bill BillData
	err = json.Unmarshal(body, &bill)
	if err != nil {
		log.Println("Error unmarshalling json:", err)
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Received Bill Data: %+v\n", bill)

	// Transaction start
	tx, err := gwdb.GDbconn.Begin()
	if err != nil {
		log.Println("Tx Start Error:", err)
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	// Insert Bill
	billQuery := `
		INSERT INTO bills (
			bill_number, bill_date,
			recipient_name, recipient_address, recipient_city, recipient_state, recipient_gstin,
			consignee_name, consignee_address, consignee_city, consignee_gstin, consignee_drug_lic_no,
			gst_invoice_no, gst_invoice_date, reverse_charge, supplier_ref_no, payment_terms, delivery_terms,
			contact_person, contact_number, email, dispatched_through, destination, remark,
			sub_total, discount_amount, grand_total,
			bank_account_number, bank_ifsc_code, bank_name, bank_branch
		) VALUES (?, NOW(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err = tx.Exec(billQuery,
		bill.BillNumber,
		bill.Recipient.Name, bill.Recipient.Address, bill.Recipient.City, bill.Recipient.State, bill.Recipient.Gstin,
		bill.Consignee.Name, bill.Consignee.Address, bill.Consignee.City, bill.Consignee.Gstin, bill.Consignee.DrugLicNo,
		bill.Invoice.GstInvoiceNo, bill.Invoice.GstInvoiceDate, bill.Invoice.ReverseCharge, bill.Invoice.SupplierRefNo,
		bill.Invoice.PaymentTerms, bill.Invoice.DeliveryTerms, bill.Invoice.ContactPerson, bill.Invoice.ContactNumber,
		bill.Invoice.Email, bill.Invoice.DispatchedThrough, bill.Invoice.Destination, bill.Invoice.Remark,
		bill.SubTotal, bill.Discount.Amount, bill.GrandTotal,
		bill.Bank.AccountNumber, bill.Bank.IfscCode, bill.Bank.BankName, bill.Bank.Branch,
	)

	if err != nil {
		tx.Rollback()
		log.Println("Insert Bill Error:", err)
		http.Error(w, "Failed to save bill: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert Bill Items
	itemQuery := `
		INSERT INTO bill_items (
			bill_number, medicine_name, batch_id, expiry_date, quantity, mrp, gst, amount
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	for _, item := range bill.BillingItems {
		_, err = tx.Exec(itemQuery,
			bill.BillNumber, item.MedicineName, item.BatchId, item.ExpiryDate,
			item.Quantity, item.Mrp, item.Gst, item.Amount,
		)
		if err != nil {
			tx.Rollback()
			log.Println("Insert Item Error:", err)
			http.Error(w, "Failed to save bill items: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Optional: Deduct stock here
		stockUpdateQuery := `UPDATE medicines SET quantity = quantity - ? WHERE name = ? AND batch_id = ?`
		_, err = tx.Exec(stockUpdateQuery, item.Quantity, item.MedicineName, item.BatchId)
		if err != nil {
			log.Println("Warning: Failed to update stock for", item.MedicineName, err)
			// Decide if this should rollback transaction or just log warning.
			// For now, let's just log it.
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("Tx Commit Error:", err)
		http.Error(w, "DB Commit Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Bill saved successfully"})
}

// GetBillDetails fetches a single bill with all its items
func GetBillDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBillDetails (+) - Request Received")
	defer log.Println("GetBillDetails (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	billNumber := r.URL.Query().Get("billNumber")
	if billNumber == "" {
		log.Println("Error: Missing billNumber parameter")
		http.Error(w, "Missing billNumber parameter", http.StatusBadRequest)
		return
	}

	log.Printf("Fetching details for Bill Number: %s\n", billNumber)

	// Fetch Bill Header
	billQuery := `
		SELECT bill_number, recipient_name, recipient_address, recipient_city, recipient_state, recipient_gstin,
		       consignee_name, consignee_address, consignee_city, consignee_gstin, consignee_drug_lic_no,
		       gst_invoice_no, gst_invoice_date, reverse_charge, supplier_ref_no, payment_terms, delivery_terms,
		       dispatched_through, destination, contact_person, contact_number, remark,
		       bank_account_number, bank_ifsc_code, bank_name, bank_branch,
		       sub_total, discount_type, discount_value, discount_amount, grand_total, party_name, created_at
		FROM bills WHERE bill_number = ?
	`
	row := gwdb.GDbconn.QueryRow(billQuery, billNumber)

	var bill BillData
	var createdAt string      // Temp storage for date
	var invoiceDateStr string // Temp storage for date

	err := row.Scan(
		&bill.BillNumber,
		&bill.Recipient.Name, &bill.Recipient.Address, &bill.Recipient.City, &bill.Recipient.State, &bill.Recipient.Gstin,
		&bill.Consignee.Name, &bill.Consignee.Address, &bill.Consignee.City, &bill.Consignee.Gstin, &bill.Consignee.DrugLicNo,
		&bill.Invoice.GstInvoiceNo, &invoiceDateStr, &bill.Invoice.ReverseCharge, &bill.Invoice.SupplierRefNo,
		&bill.Invoice.PaymentTerms, &bill.Invoice.DeliveryTerms, &bill.Invoice.DispatchedThrough, &bill.Invoice.Destination,
		&bill.Invoice.ContactPerson, &bill.Invoice.ContactNumber, &bill.Invoice.Remark,
		&bill.Bank.AccountNumber, &bill.Bank.IfscCode, &bill.Bank.BankName, &bill.Bank.Branch,
		&bill.SubTotal, &bill.Discount.Type, &bill.Discount.Value, &bill.Discount.Amount, &bill.GrandTotal, &bill.PartyName,
		&createdAt,
	)

	if err != nil {
		log.Println("Error fetching bill header:", err)
		http.Error(w, "Bill not found", http.StatusNotFound)
		return
	}
	bill.Invoice.GstInvoiceDate = invoiceDateStr // Assign string date

	// Fetch Bill Items
	itemsQuery := `
		SELECT medicine_name, batch_id, expiry_date, quantity, mrp, gst, amount
		FROM bill_items WHERE bill_number = ?
	`
	rows, err := gwdb.GDbconn.Query(itemsQuery, billNumber)
	if err != nil {
		log.Println("Error fetching bill items:", err)
		http.Error(w, "Failed to fetch bill items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item BillingItem
		if err := rows.Scan(&item.MedicineName, &item.BatchId, &item.ExpiryDate, &item.Quantity, &item.Mrp, &item.Gst, &item.Amount); err != nil {
			log.Println("Error scanning bill item:", err)
			continue
		}
		bill.BillingItems = append(bill.BillingItems, item)
	}

	log.Printf("Successfully fetched bill %s with %d items\n", billNumber, len(bill.BillingItems))
	json.NewEncoder(w).Encode(bill)
}

func GetBills(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBills (+) - Fetching History Summary")
	defer log.Println("GetBills (-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	query := `
		SELECT bill_number, created_at, party_name, grand_total 
		FROM bills 
		ORDER BY created_at DESC
	`

	log.Println("Executing Query:", query)

	rows, err := gwdb.GDbconn.Query(query)
	if err != nil {
		log.Println("DB Error executing GetBills query:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type BillSummary struct {
		BillNumber string  `json:"bill_number"`
		Date       string  `json:"date"`
		Total      float64 `json:"total"`
		Customer   string  `json:"customer"`
	}

	var bills []BillSummary
	for rows.Next() {
		var b BillSummary
		var dateStr string
		if err := rows.Scan(&b.BillNumber, &dateStr, &b.Customer, &b.Total); err != nil {
			log.Println("Error scanning row in GetBills:", err)
			continue
		}
		b.Date = dateStr
		bills = append(bills, b)
	}

	log.Printf("Found %d bills in history\n", len(bills))
	json.NewEncoder(w).Encode(bills)
}

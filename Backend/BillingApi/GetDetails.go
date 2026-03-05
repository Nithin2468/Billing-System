package billingapi

import (
	"billingapp/gwdb"
	"encoding/json"
	"log"
	"net/http"
)

type UserManagement struct {
	User_id    string `json:"user_id"`
	UserName   string `json:"user_name"`
	UserRole   string `json:"user_role"`
	Acc_status string `json:"account_status"`
}

type UpdateUserReq struct {
	Userid    string `json:"user_id"`
	AccStatus bool   `json:"status"` // true = ACTIVE, false = INACTIVE
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUser(+)")
	defer log.Println("GetUser(-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	query := `
		SELECT ud.user_id, ud.user_name, ld.user_role, ld.account_status
		FROM login_deatils ld 
		JOIN user_details ud ON ld.user_id = ud.user_id
	`
	rows, err := gwdb.GDbconn.Query(query)
	if err != nil {
		log.Println("Unable to fetch details:", err)
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []UserManagement
	for rows.Next() {
		var u UserManagement
		if err := rows.Scan(&u.User_id, &u.UserName, &u.UserRole, &u.Acc_status); err != nil {
			log.Println("Error scanning user:", err)
			continue
		}
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

func UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateUserStatus(+)")
	defer log.Println("UpdateUserStatus(-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req UpdateUserReq
	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	newStatus := "INACTIVE"
	if req.AccStatus {
		newStatus = "ACTIVE"
	}

	query := `UPDATE login_deatils SET account_status = ? WHERE user_id = ?`
	_, err := gwdb.GDbconn.Exec(query, newStatus, req.Userid)
	if err != nil {
		log.Println("Error updating user status:", err)
		http.Error(w, "DB Update Failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "User status updated successfully",
	})
}

package billingapi

import (
	"billingapp/gwdb"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserReq struct {
	Userid     string `json:"user_id"`
	UserEmail  string `json:"user_email"`
	UserMoblie string `json:"user_phone"`
	Password   string `json:"password"`
	UserName   string `json:"user_name"`
	UserRole   string `json:"user_role"`
	CreatedBy  string `json:"created_by"`
}

type ApiResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func CreateUserApi(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateUserApi(+) - Request Received")
	defer log.Println("CreateUserApi(-)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Println("req body ", string(body))
	var req CreateUserReq
	if err := json.Unmarshal(body, &req); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Creating User: %s (Role: %s)\n", req.Userid, req.UserRole)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, "Password encryption failed", http.StatusInternalServerError)
		return
	}

	query := `
		INSERT INTO login_deatils 
		(user_id, password_hash, user_role, created_by, updated_by)
		VALUES (?, ?, ?, "Admin", "Admin")
	`

	_, err = gwdb.GDbconn.Exec(query, req.Userid, string(hashedPass), req.UserRole)
	if err != nil {
		log.Println("DB Error inserting user:", err)
		json.NewEncoder(w).Encode(ApiResponse{
			Message: "User already exists or DB error",
			Status:  false,
		})
		return
	}
	err = InsertLoginDetails(req)
	if err != nil {
		log.Println("DB Error inserting user details:", err)
		json.NewEncoder(w).Encode(ApiResponse{
			Message: "User already exists or DB error",
			Status:  false,
		})
		return
	}
	log.Println("User created successfully")
	json.NewEncoder(w).Encode(ApiResponse{
		Message: "User created successfully",
		Status:  true,
	})

	log.Println("CreateUserApi(-)")
}

type LogHistoryReq struct {
	Userid    string `json:"user_id"`
	Username  string `json:"user_name"`
	UserRole  string `json:"user_role"`
	CreatedBy string `json:"created_by"`
}

func InsertLoginDetails(pDetails CreateUserReq) error {
	query := `
		INSERT INTO user_details 
		(user_id, user_name, user_role, user_phone,user_email,created_by, updated_by)
		VALUES (?, ?, ?, ?,?, "Admin", "Admin")
	`
	_, err := gwdb.GDbconn.Exec(query, pDetails.Userid, pDetails.UserName, pDetails.UserRole, pDetails.UserMoblie, pDetails.UserEmail)
	if err != nil {
		log.Println("DB Error inserting user details:", err)
		return err
	}
	return nil
}

package billingapi

import (
	"billingapp/gwdb"
	"billingapp/middleware"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginDetails struct {
	Userid   string `json:"user_id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Details UserDetails `json:"details"`
}
type UserDetails struct {
	UserName    string `json:"user_name"`
	Mobileno    string `json:"user_mobile"`
	Emailaddrss string `json:"user_email"`
	User_role   string `json:"user_role"`
}

func LoginApi(w http.ResponseWriter, r *http.Request) {
	log.Println("Login(+)")
	defer log.Println("Login(-)")

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
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req LoginDetails
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	query := `
		SELECT password_hash, account_status, user_role
		FROM login_deatils 
		WHERE user_id='` + req.Userid + `'
	`
	log.Println("query", query)
	var lPass, lAcc_Status, lRole string
	err = gwdb.GDbconn.QueryRow(query).Scan(&lPass, &lAcc_Status, &lRole)
	if err != nil {
		log.Println("User not found 1:", err)
		json.NewEncoder(w).Encode(LoginResponse{
			Message: "Invalid credentials",
			Status:  false,
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(lPass), []byte(req.Password))
	if err != nil {
		json.NewEncoder(w).Encode(LoginResponse{
			Message: "Invalid credentials",
			Status:  false,
		})
		return
	}

	if lAcc_Status != "ACTIVE" {
		json.NewEncoder(w).Encode(LoginResponse{
			Message: "Account Inactive",
			Status:  false,
		})
		return
	}
	lDeatails, lErr := GetUSerData(req.Userid)
	if lErr != nil {
		log.Println("Error while fetching Data ", lErr)
	}
	// Fallback to role from login_deatils if not found in user_details or if GetUSerData failed
	if lDeatails.User_role == "" {
		lDeatails.User_role = lRole
	}

	token, err := middleware.GenerateJWT(req.Userid)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   false, // true in prod (HTTPS)
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   1800,
	})
	lErr = InsertLoginHistory(req.Userid, lDeatails.UserName, lDeatails.User_role)
	if lErr != nil {
		log.Println("Error while inserting Login History ", lErr)
	}
	json.NewEncoder(w).Encode(LoginResponse{
		Message: "Login successful",
		Status:  true,
		Details: lDeatails,
	})

}

func InsertLoginHistory(pUserid string, pUsername string, pUserrole string) error {
	log.Println("log_history+")
	defer log.Println("Log history -")
	query := `
		INSERT INTO log_history
		(user_id, user_name, user_role, login_time, created_by, updated_by)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := gwdb.GDbconn.Exec(query, pUserid, pUsername, pUserrole, time.Now(), "Admin", "Admin")
	if err != nil {
		log.Println("DB Error:", err)
		return err
	}
	return nil
}
func GetUSerData(pUserid string) (UserDetails, error) {
	var lClientDetails UserDetails
	log.Println("user_id", pUserid)
	query := `
		SELECT user_name ,user_phone,user_email,user_role
		FROM user_details 
		WHERE user_id='` + pUserid + `'
	`
	lErr := gwdb.GDbconn.QueryRow(query).Scan(&lClientDetails.UserName, &lClientDetails.Mobileno, &lClientDetails.Emailaddrss, &lClientDetails.User_role)
	if lErr != nil {
		log.Println("User Detais not found:", lErr)
		return lClientDetails, lErr
	}

	return lClientDetails, nil
}

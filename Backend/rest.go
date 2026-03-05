package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my_secret_key")

// ====== LOGIN REQUEST ======
type LoginReq struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

// ====== SESSION LOGIN ======
func sessionLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginReq
	json.NewDecoder(r.Body).Decode(&req)

	if req.User != "abc" || req.Pass != "1234" {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "SESSION_ABC_123",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   10,
	})

	w.Write([]byte("Session login success"))
}

// ====== SESSION PROTECTED ======
func sessionDashboard(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value != "SESSION_ABC_123" {
		http.Error(w, "Unauthorized (session)", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Welcome (SESSION AUTH)"))
}

// ====== JWT LOGIN ======
func jwtLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginReq
	json.NewDecoder(r.Body).Decode(&req)

	if req.User != "abc" || req.Pass != "1234" {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": req.User,
		"exp":  time.Now().Add(10 * time.Minute).Unix(),
	})

	tokenStr, _ := token.SignedString(jwtSecret)

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenStr,
	})
}

// ====== JWT PROTECTED ======
func jwtDashboard(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	tokenStr := auth[len("Bearer "):]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Welcome (JWT AUTH)"))
}

func restMain() {
	http.HandleFunc("/login/session", sessionLogin)
	http.HandleFunc("/dashboard/session", sessionDashboard)

	http.HandleFunc("/login/jwt", jwtLogin)
	http.HandleFunc("/dashboard/jwt", jwtDashboard)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

// func main() {
// str1 := "Abbi2 Meet5 yo6u 3nice 0hi to4 dea1r"

// var all []map[int]string
// var akey []int
// vals := strings.Split(str1, " ")
// for _, val := range vals {
// 	var m1 = make(map[int]string)
// 	var value string
// 	var key int
// 	for _, i := range val {
// 		if i >= '0' && i <= '9' {
// 			key, _ = strconv.Atoi(string(i))
// 		} else {
// 			value += string(i)
// 		}
// 	}
// 	akey = append(akey, key)
// 	m1[key] = value
// 	all = append(all, m1)
// }
// var final string
// for k1, _ := range all {
// 	for _,val3:=range all{
// 		for p, val4 := range val3 {
// 			if k1==p{
// 				final +=val4+" "
// 			}
// 		}
// 	}

// }

// fmt.Println("MAp:", final)

// }

package main

import (
	billingapi "billingapp/BillingApi"
	commonapi "billingapp/CommonAPi"
	"billingapp/gwdb"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02_15-04-05")+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Errored in log Creation ", err.Error())
	} else {
		defer file.Close()
		log.SetOutput(file)
		fmt.Println("Log File Created succesfully")
	}
	var lerr error
	gwdb.GDbconn, lerr = gwdb.DbConnection()
	if lerr != nil {
		log.Println("lERrr", lerr.Error())
	}
	defer gwdb.GDbconn.Close()
	// log.Println("DB Host:", cfg.Host)
	http.HandleFunc("/billingrecords", billingapi.BillDetails)
	http.HandleFunc("/downloadpdff", billingapi.DownloadPdf)
	http.HandleFunc("/get-bank-details", commonapi.GetBankDetails)
	// http.HandleFunc("/middleware", middleware.)

	http.HandleFunc("/loginApi", billingapi.LoginApi)     //DONE
	http.HandleFunc("/newuser", billingapi.CreateUserApi) //DONE

	// New Routes
	http.HandleFunc("/add-medicine", billingapi.AddMedicine)
	http.HandleFunc("/get-medicines", billingapi.GetMedicines)
	http.HandleFunc("/update-medicine", billingapi.UpdateMedicine)
	http.HandleFunc("/get-bills", billingapi.GetBills)
	http.HandleFunc("/get-bill-details", billingapi.GetBillDetails)
	http.HandleFunc("/reports/stock-sales", billingapi.GetStockReport)

	// User Management
	http.HandleFunc("/get-users", billingapi.GetUser)
	http.HandleFunc("/update-user-status", billingapi.UpdateUserStatus)

	fmt.Println("http://localhost:8090")
	http.ListenAndServe(":8090", nil)

}

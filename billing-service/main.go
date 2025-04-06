// ========== billing-service/main.go ==========
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func processBilling(w http.ResponseWriter, r *http.Request) {
	log.Println("Billing processed")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "billing_success"})
}

func refundBilling(w http.ResponseWriter, r *http.Request) {
	log.Println("Billing refunded")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "billing_refunded"})
}

func main() {
	http.HandleFunc("/pay", processBilling)
	http.HandleFunc("/refund", refundBilling)
	log.Println("Payment Service running on port 8002")
	http.ListenAndServe(":8002", nil)
}

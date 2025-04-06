// ========== delivery-service/main.go ==========
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func deliveryOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Delivery process started")
	if rand.Intn(2) == 0 {
		log.Println("Delivery failed")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "delivery_failed"})
		return
	}
	log.Println("Delivery successful")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "delivery_success"})
}

func cancelDelivery(w http.ResponseWriter, r *http.Request) {
	log.Println("Delivery cancelled")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "delivery_cancelled"})
}

func main() {
	http.HandleFunc("/delivery", deliveryOrder)
	http.HandleFunc("/cancel", cancelDelivery)
	log.Println("Delivery Service running on port 8003")
	http.ListenAndServe(":8003", nil)
}

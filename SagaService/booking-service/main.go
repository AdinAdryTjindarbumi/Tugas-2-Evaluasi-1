// ========== booking-service/main.go ==========
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

func createBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("Booking created")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "booking_created"})
}

func cancelBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("Booking cancelled")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "booking_cancelled"})
}

func main() {
	http.HandleFunc("/create", createBooking)
	http.HandleFunc("/cancel", cancelBooking)
	log.Println("Order Service running on port 8001")
	http.ListenAndServe(":8001", nil)
}

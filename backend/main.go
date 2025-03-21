package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	DB, err = InitDB() // Используем экспортированную функцию InitDB
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer DB.Close()

	if err := SetupDatabase(DB); err != nil {
		log.Fatalf("Error setting up database schema: %v", err)
	}

	if err := InsertSampleData(DB); err != nil {
		log.Fatalf("Error inserting sample data: %v", err)
	}
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/api/hotels", GetHotelsHandler)
	http.HandleFunc("/api/hotels/create", CreateHotelHandler)
	http.HandleFunc("/health", HealthCheckHandler)
	http.HandleFunc("/api/rooms", GetRoomsHandler)
	http.HandleFunc("/api/customers", GetCustomersHandler)
	http.HandleFunc("/api/bookings", GetBookingsHandler)
	http.HandleFunc("/api/reviews", GetReviewsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию
	}

	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, EnableCORS(http.DefaultServeMux)))
}

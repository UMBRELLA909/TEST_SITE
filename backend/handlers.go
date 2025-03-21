package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// handlers.go

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Отправляем HTML-файл главной страницы
	http.ServeFile(w, r, "./frontend/public/index.html")
}

func GetRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, hotel_id, category, capacity, price, available FROM rooms")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rooms []Room
	for rows.Next() {
		var room Room
		if err := rows.Scan(&room.ID, &room.HotelID, &room.Category, &room.Capacity, &room.Price, &room.Available); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(rooms); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCustomersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, first_name, last_name, email, phone_number, loyalty_points FROM customers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.PhoneNumber, &customer.LoyaltyPoints); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(customers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetBookingsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, hotel_id, room_id, customer_id, check_in_date, check_out_date, total_price, booking_date, cancelled, cancellation_fee FROM bookings")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.ID, &booking.HotelID, &booking.RoomID, &booking.CustomerID, &booking.CheckInDate, &booking.CheckOutDate, &booking.TotalPrice, &booking.BookingDate, &booking.Cancelled, &booking.CancellationFee); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bookings = append(bookings, booking)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetReviewsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, hotel_id, customer_id, rating, comment, review_date FROM reviews")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		if err := rows.Scan(&review.ID, &review.HotelID, &review.CustomerID, &review.Rating, &review.Comment, &review.ReviewDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reviews); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetHotelsHandler(w http.ResponseWriter, r *http.Request) {
	hotels, err := GetHotelsFromDB() // Use exported GetHotelsFromDB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(hotels); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateHotelHandler(w http.ResponseWriter, r *http.Request) {
	var hotel Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := CreateHotel(hotel); err != nil { // Use exported CreateHotel
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Hotel created successfully!")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if CheckDBConnection() { // Use exported CheckDBConnection
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Backend and database are healthy")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Backend is healthy, but database connection failed")
	}
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

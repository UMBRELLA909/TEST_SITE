package main

import "time"

type Hotel struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"`
	Rooms []Room `json:"rooms"`
}

type Room struct {
	ID        int     `json:"id"`
	HotelID   int     `json:"hotel_id"`
	Category  string  `json:"category"`  // "budget", "standard", "suite", "luxury"
	Capacity  int     `json:"capacity"`  // 1 (single), 2 (double)
	Price     float64 `json:"price"`     // Price per night
	Available bool    `json:"available"` // Indicates if the room is currently available
}

type Booking struct {
	ID              int       `json:"id"`
	HotelID         int       `json:"hotel_id"`
	RoomID          int       `json:"room_id"`
	CustomerID      int       `json:"customer_id"` //  (User)
	CheckInDate     time.Time `json:"check_in_date"`
	CheckOutDate    time.Time `json:"check_out_date"`
	TotalPrice      float64   `json:"total_price"`
	BookingDate     time.Time `json:"booking_date"`
	Cancelled       bool      `json:"cancelled"`
	CancellationFee float64   `json:"cancellation_fee"`
}

type Customer struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	LoyaltyPoints int    `json:"loyalty_points"`
}

type Review struct {
	ID         int       `json:"id"`
	HotelID    int       `json:"hotel_id"`
	CustomerID int       `json:"customer_id"`
	Rating     int       `json:"rating"` // Scale from 1 to 5
	Comment    string    `json:"comment"`
	ReviewDate time.Time `json:"review_date"`
}

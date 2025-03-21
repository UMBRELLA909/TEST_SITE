package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB // Экспортируем переменную DB

func InitDB() (*sql.DB, error) {
	// Load environment variables
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost" // Default value if not set
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432" // Default value if not set
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres" // Default value if not set
	}
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		dbPass = "123" // Default value if not set
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "hotel_bd" // Default value if not set
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Successfully connected to the database!")
	return DB, nil
}

func GetHotelsFromDB() ([]Hotel, error) {
	rows, err := DB.Query("SELECT id, name, city FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []Hotel
	for rows.Next() {
		var hotel Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.City); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return hotels, nil
}

func CreateHotel(hotel Hotel) error {
	_, err := DB.Exec("INSERT INTO hotels (name, city) VALUES ($1, $2)", hotel.Name, hotel.City)
	if err != nil {
		return err
	}
	return nil
}

func CheckDBConnection() bool {
	err := DB.Ping()
	if err != nil {
		log.Println("Database connection failed:", err)
		return false
	}
	return true
}

func SetupDatabase(db *sql.DB) error {
	// Create hotels table
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS hotels (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            city VARCHAR(255) NOT NULL
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create hotels table: %w", err)
	}

	// Create rooms table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS rooms (
            id SERIAL PRIMARY KEY,
            hotel_id INT REFERENCES hotels(id) ON DELETE CASCADE,
            category VARCHAR(50) NOT NULL,
            capacity INT NOT NULL,
            price DECIMAL(10, 2) NOT NULL,
            available BOOLEAN NOT NULL DEFAULT TRUE
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create rooms table: %w", err)
	}

	// Create customers table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS customers (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(255) NOT NULL,
            last_name VARCHAR(255) NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            phone_number VARCHAR(20),
            loyalty_points INT DEFAULT 0
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create customers table: %w", err)
	}

	// Create bookings table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS bookings (
            id SERIAL PRIMARY KEY,
            hotel_id INT REFERENCES hotels(id) ON DELETE CASCADE,
            room_id INT REFERENCES rooms(id) ON DELETE CASCADE,
            customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
            check_in_date DATE NOT NULL,
            check_out_date DATE NOT NULL,
            total_price DECIMAL(10, 2) NOT NULL,
            booking_date TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
            cancelled BOOLEAN DEFAULT FALSE,
            cancellation_fee DECIMAL(10, 2) DEFAULT 0.00
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create bookings table: %w", err)
	}

	// Create reviews table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS reviews (
            id SERIAL PRIMARY KEY,
            hotel_id INT REFERENCES hotels(id) ON DELETE CASCADE,
            customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
            rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
            comment TEXT,
            review_date TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc')
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create reviews table: %w", err)
	}

	fmt.Println("Database schema created successfully!")
	return nil
}

func InsertSampleData(db *sql.DB) error {
	// Очистка всех таблиц перед вставкой данных
	tables := []string{"bookings", "reviews", "rooms", "customers", "hotels"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s;", table))
		if err != nil {
			return fmt.Errorf("failed to clear table %s: %w", table, err)
		}
	}

	// Вставка тестовых отелей
	_, err := db.Exec(`
        INSERT INTO hotels (name, city) VALUES
        ('Grand Hotel', 'New York'),
        ('Sunset Resort', 'Los Angeles'),
        ('Mountain View Lodge', 'Aspen');
    `)
	if err != nil {
		return fmt.Errorf("failed to insert sample hotels: %w", err)
	}

	// Вставка тестовых клиентов
	_, err = db.Exec(`
        INSERT INTO customers (first_name, last_name, email, phone_number) VALUES
        ('John', 'Doe', 'john.doe@example.com', '123-456-7890'),
        ('Jane', 'Smith', 'jane.smith@example.com', '987-654-3210')
        ON CONFLICT (email) DO UPDATE SET
            first_name = EXCLUDED.first_name,
            last_name = EXCLUDED.last_name,
            phone_number = EXCLUDED.phone_number;
    `)
	if err != nil {
		return fmt.Errorf("failed to insert sample customers: %w", err)
	}

	// Вставка тестовых комнат
	_, err = db.Exec(`
        INSERT INTO rooms (hotel_id, category, capacity, price) VALUES
        (1, 'standard', 2, 150.00),
        (1, 'suite', 2, 250.00),
        (1, 'luxury', 2, 400.00),
        (2, 'standard', 2, 120.00),
        (2, 'suite', 2, 220.00),
        (2, 'luxury', 2, 350.00),
        (3, 'standard', 2, 180.00),
        (3, 'suite', 2, 300.00),
        (3, 'luxury', 2, 450.00);
    `)
	if err != nil {
		return fmt.Errorf("failed to insert sample rooms: %w", err)
	}

	fmt.Println("Sample data inserted successfully!")
	return nil
}

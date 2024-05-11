package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func migrate(db *sql.DB) {
	query := `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		age INTEGER,
		gender VARCHAR(10),
		password TEXT NOT NULL,
		is_verified BOOLEAN DEFAULT false,
		created_at TIMESTAMP ,
		updated_at TIMESTAMP 
	);

	CREATE TABLE IF NOT EXISTS account_payments (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name TEXT NOT NULL,
		type varchar(20) NOT NULL,
		user_id UUID NOT NULL,
		amount real, -- Assumes two decimal places for currency amounts
		created_at TIMESTAMP ,
		updated_at TIMESTAMP 
	);

	CREATE TABLE IF NOT EXISTS historical_transactions (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		account_payment_id_payed UUID,  -- Nullable
		account_payment_id_received UUID,  -- Nullable
		total_transaction real,  -- Adjust size and precision as necessary
		currency VARCHAR(10) NOT NULL,
		status_transaction VARCHAR(50) NOT NULL,
		type_transaction VARCHAR(50) NOT NULL,
		created_at TIMESTAMP ,
		updated_at TIMESTAMP 
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// do migrasi
	migrate(db)
	fmt.Println("succes migrate db")
}

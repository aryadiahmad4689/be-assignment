package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitPostgres(driver string, source string) (*sql.DB, error) {
	// Membuka koneksi ke database SQLite
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}

	// Menguji koneksi database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

package postgres

import (
	"be-assesment/src/app/db"
	"context"
	"database/sql"
	"fmt"
	"os"
)

func InitPostgresMaster(ctx context.Context) *sql.DB {

	clientConnect, err := db.Init(ctx, db.Config{
		Driver: os.Getenv("DB_DRIVER"),
		Source: os.Getenv("DB_SOURCE"),
	})
	if err != nil {
		fmt.Printf("cannot connect to postgress: %s\n", err.Error())
		os.Exit(1)
	}
	return clientConnect.(*sql.DB)
}
func InitPostgresSlave(ctx context.Context) *sql.DB {
	clientConnect, err := db.Init(ctx, db.Config{
		Driver: os.Getenv("DB_DRIVER"),
		Source: os.Getenv("DB_SOURCE"),
	})
	if err != nil {
		fmt.Printf("cannot connect to postgress: %s\n", err.Error())
		os.Exit(1)
	}
	return clientConnect.(*sql.DB)
}

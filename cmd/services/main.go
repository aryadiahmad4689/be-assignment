package main

import (
	"fmt"
	"log"
	"os"

	"be-assesment/src/app"
	"be-assesment/src/modules"

	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var (
		app    = app.Init()
		router = app.GetHttpRouter()
		module = modules.Init(app)
	)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://", "http://"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	router.Mount("/v1/payment", module.PaymentManager.GetHttpRouter())
	router.Mount("/v1/account", module.AccountManager.GetHttpRouter())

	err = app.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

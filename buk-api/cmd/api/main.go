package main

import (
	"buk/buk-backend-service/internal/appointment"
	postgres "buk/buk-backend-service/internal/db"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	log.Println("Starting the application...")

	db := postgres.ConnectToDB(context.Background())
	defer db.Close()

	r := gin.Default()

	importRoutes(r, db)

	err := r.Run()
	if err != nil {
		log.Fatalln("Error while starting the service API:", err)
	}

	log.Println("Application Closed !")
}

func importRoutes(r *gin.Engine, json *sqlx.DB) {
	appointment.Routes(r, json)
}

package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"schedule-ms/models"
)

//var DB *gorm.DB

type Handler struct {
	DB *gorm.DB
}

func Init() Handler {
	host := "schedule-db"
	port := "5432"
	dbName := "scheduler_svc"
	dbUser := "postgres"
	password := "password123"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(models.Appointment{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Scheduler database connection successful...")

	return Handler{db}
}

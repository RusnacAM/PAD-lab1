package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"staff-ms/models"
)

type Handler struct {
	DB *gorm.DB
}

func Init(host string, dbName string) Handler {
	port := "5432"
	dbUser := "postgres"
	password := os.Getenv("STAFF_DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(models.StaffRecord{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Staff records database connection successful...")

	return Handler{db}
}

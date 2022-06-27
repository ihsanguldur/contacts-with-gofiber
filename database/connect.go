package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rehber/config"
	"rehber/models"
	"strconv"
)

//Connect open a new database connection.
//This function will migrate models.User and models.PhoneNumber structs.
func Connect() {
	var err error
	p := config.Config("DB_PORT")
	port, _ := strconv.Atoi(p)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error While Database Connection.")
	}
	fmt.Println("Connection Opened to Database.")

	if err = DB.AutoMigrate(&models.User{}, &models.PhoneNumber{}); err != nil {
		log.Fatal("Error While Migrating.")
	}
	fmt.Println("Database migrated.")
}

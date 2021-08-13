package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ElizeuS/gouser/database/migrations"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	str := fmt.Sprintf("host=%s user=%s  password=%s  dbname=%s  port=%s  sslmode=disable",
		goDotEnvVariable("DB_HOST"), goDotEnvVariable("DB_USER"), goDotEnvVariable("DB_PASSWORD"),
		goDotEnvVariable("DB_NAME"), goDotEnvVariable("DB_PORT"))

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)

}

func GetDaabase() *gorm.DB {
	return db
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

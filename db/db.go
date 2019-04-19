package db

import (
	"github.com/AnthonyLaiuppa/ezctf-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to postgres database and
// migrates any new models
func Init() {

	//getEnv("PG_CSTRING", "postgres://DEFAULT:SETTING@127.0.0.1:5432/dbname?connect_timeout=10")
	dbcon := os.Getenv("DBCONNSTRING")
	db, err = gorm.Open("postgres", dbcon)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	if !db.HasTable(&models.Challenge{}) {
		err := db.CreateTable(&models.Challenge{})
		if err != nil {
			log.Println("Table already exists")
		}
	}
	if !db.HasTable(&models.User{}) {
		err := db.CreateTable(&models.User{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&models.Challenge{})
	db.AutoMigrate(&models.User{})

}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

//CloseDB ... add trigger for server graceful stop
func CloseDB() {
	db.Close()
}

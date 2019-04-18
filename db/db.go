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
	db.Create(&models.User{UserName: "floridamane", Password: "testingstuff", Email: "registereduser@email.com", Score: 0, Solves: "uuid-uuid-uuid-uuid", Elevated: false})
	db.Create(&models.Challenge{Name: "babys first stack", Category: "reverse", Solves: 10, Points: 150, Author: "FloridaMane", RawText: "This is challenge 1.",Flag:"ezctf{flag1}"})
	db.Create(&models.Challenge{Name: "reverse2", Category: "reverse", Solves: 10, Points: 200, Author: "FloridaMane", RawText: "This is challenge 2.",Flag:"ezctf{flag2}"})
	db.Create(&models.Challenge{Name: "reverse3", Category: "reverse", Solves: 16, Points: 250, Author: "FloridaMane", RawText: "This is challenge 3.",Flag:"ezctf{flag3}"})
	db.Create(&models.Challenge{Name: "pwn1", Category: "pwn", Solves: 14, Points: 300, Author: "FloridaMane", RawText: "This is challenge 4.",Flag:"ezctf{flag4}"})
	db.Create(&models.Challenge{Name: "pwn2", Category: "pwn", Solves: 58, Points: 350, Author: "FloridaMane", RawText: "This is challenge 5.",Flag:"ezctf{flag5}"})
	db.Create(&models.Challenge{Name: "forensics1", Category: "forensics", Solves: 2334, Points: 500, Author: "FloridaMane", RawText: "This is challenge 6.",Flag:"ezctf{flag6}"})
	db.Create(&models.Challenge{Name: "forensics2", Category: "forsenics", Solves: 430, Points: 1000, Author: "FloridaMane", RawText: "This is challenge 7.",Flag:"ezctf{flag7}"})
	db.Create(&models.Challenge{Name: "misc1", Category: "misc", Solves: 70, Points: 10000, Author: "FloridaMane", RawText: "This is challenge 8.",Flag:"ezctf{flag8}"})
	db.Create(&models.Challenge{Name: "misc2", Category: "misc", Solves: 6, Points: 50, Author: "FloridaMane", RawText: "This is challenge 9.",Flag:"ezctf{flag9}"})
	db.Create(&models.Challenge{Name: "web1", Category: "web", Solves: 0, Points: 10, Author: "FloridaMane", RawText: "This is challenge 10.",Flag:"ezctf{flag10}"})
	db.Create(&models.Challenge{Name: "web2", Category: "web", Solves: 1, Points: 1, Author: "FloridaMane", RawText: "This is challenge 11.",Flag:"ezctf{flag11}"})

}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

//CloseDB ... add trigger for server graceful stop
func CloseDB() {
	db.Close()
}

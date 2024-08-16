package configDB

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Connect() {
	var err error
	connStr := "host=localhost user=postgres dbname=postgres sslmode=disable password=senha"
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		os.Exit(1)
	}
	fmt.Println("Database connection established")
}

func GetDB() *gorm.DB {
	return DB
}

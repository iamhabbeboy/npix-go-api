package database

import (
	"os"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Connect to mysql
func Connect() *gorm.DB {
	host, hostErr := os.LookupEnv("DB_HOST")
	name, nameErr := os.LookupEnv("DB_NAME")
	pass, passErr := os.LookupEnv("DB_PASS")
	user, userErr := os.LookupEnv("DB_USER")

	if !hostErr {
		log.Fatal("DB_HOST data not defined in .env")
	}

	if !nameErr {
		log.Fatal("DB_NAME data not defined in .env")
	}
	
	if !passErr {
		log.Fatal("DB_PASS not defined in .env")
	} 
	
	if !userErr {
		log.Fatal("DB_USER data not defined in .env")
	}

	credentials := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, name)

	db, err := gorm.Open("mysql", credentials)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
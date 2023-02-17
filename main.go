package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type User struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
}

type UserBody struct {
	name  string
	email string
}

func connect() {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	var err error

	db, err = gorm.Open(dialect, dbURL)
	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("Connection success")
	}

	defer db.Close()

}

func main() {
	connect()
}

func do_migrations() {
	db.AutoMigrate(&User{})
}

func insert(userBody []UserBody) {
	for i := range userBody {
		user := User{Name: userBody[i].name, Email: userBody[i].email}
		db.Create(user)
	}
}

func get() {

}

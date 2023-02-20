package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

}
func get_users(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)

}

func get_user(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User

	db.First(&user, params["id"])

	json.NewEncoder(w).Encode(user)
}

func create_user(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	newUser := db.Create(&user)

	json.NewEncoder(w).Encode(newUser)
}

func delete_user(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	db.Delete(&user)
	json.NewEncoder(w).Encode(user)
}

func update_user(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	db.Update(&user, params["id"])

	json.NewEncoder(w).Encode(user)
}

func do_migrations(w http.ResponseWriter, r *http.Request) {
	db.AutoMigrate(&User{})
}

func main() {
	connect()

	router := mux.NewRouter()

	router.HandleFunc("/users", get_users).Methods("GET")
	router.HandleFunc("/users", create_user).Methods("POST")
	router.HandleFunc("/users/{id}", delete_user).Methods("DELETE")
	router.HandleFunc("/users/{id}", update_user).Methods("PUT")
	router.HandleFunc("/migrate", do_migrations).Methods("GET")
	http.ListenAndServe(":8007", router)
	defer db.Close()

}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var hand string
var path = "templates/"

func main() {

	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/register", register)
	// http.HandleFunc("/confirmation", confirmation)
	fmt.Println("Server is running.....")
	http.ListenAndServe(":8000", nil)

}

func register(response http.ResponseWriter, request *http.Request) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var firstname = request.FormValue("firstname")
	var lastname = request.FormValue("lastname")
	var street = request.FormValue("street")
	var city = request.FormValue("city")
	var state = request.FormValue("state")
	var zip = request.FormValue("zip")
	var email = request.FormValue("email")
	var username = request.FormValue("username")
	var pass = request.FormValue("pass")

	// fmt.Println("First name", firstname)
	// fmt.Println("Last name", lastname)
	// fmt.Println("Address", street)
	// fmt.Println("City", city)
	// fmt.Println("State", state)
	// fmt.Println("Zip Code", zip)
	// fmt.Println("Email", email)
	// fmt.Println("Username", username)
	// fmt.Println("Password", pass)

	result, err := db.Exec("insert into client (firstname, lastname, street, city, statec, zip, email, username, pass) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)", firstname, lastname, street, city, state, zip, email, username, pass)
	if err != nil {
		log.Fatal(err)
	}
	temp, err := template.ParseFiles("templates/confirmation.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(response, nil)
	fmt.Println(result)
}

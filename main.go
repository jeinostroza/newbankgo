package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/190930-UTA-CW-Go/newbankgo/functions"
	_ "github.com/lib/pq"
)

//Client is a structure of client
// type Client struct {
// 	firstname string `json:firstname`
// 	lastname  string `json:lastname`
// 	street    string `json:street`
// 	city      string `json:city`
// 	state     string `json:state`
// 	zip       string `json:zip`
// 	email     string `json:email`
// 	username  string `json:username`
// 	pass      string "json:pass"
// }

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// var hand string
// var path = "templates/"
var usernameClient string
var passClient string

func main() {

	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/open", open)
	http.HandleFunc("/deposit", deposit)
	fmt.Println("Server is running.....")
	http.ListenAndServe(":8000", nil)

}

func deposit(w http.ResponseWriter, r *http.Request) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var deposit = r.FormValue("deposit")
	dep, _ := strconv.ParseFloat(deposit, 32)

	fmt.Println(dep)
}

func open(w http.ResponseWriter, r *http.Request) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var monthlyIncomes = r.FormValue("monthlyIncomes")
	var monthlyExpenses = r.FormValue("monthlyExpenses")

	mi, _ := strconv.ParseFloat(monthlyIncomes, 32)
	me, _ := strconv.ParseFloat(monthlyExpenses, 32)

	temp, err := template.ParseFiles("templates/openconfir.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)

	usernameSearch := usernameClient

	functions.UpdateIncomes(db, mi, usernameSearch)
	functions.UpdateExpenses(db, me, usernameSearch)
}

func register(w http.ResponseWriter, r *http.Request) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var firstname = r.FormValue("firstname")
	var lastname = r.FormValue("lastname")
	var street = r.FormValue("street")
	var city = r.FormValue("city")
	var state = r.FormValue("state")
	var zip = r.FormValue("zip")
	var email = r.FormValue("email")
	var username = r.FormValue("username")
	var pass = r.FormValue("pass")

	// client := Client{
	// 	firstname: firstname,
	// 	lastname:  lastname,
	// 	street:    street,
	// 	city:      city,
	// 	state:     state,
	// 	zip:       zip,
	// 	email:     email,
	// 	username:  username,
	// }

	// fmt.Println(client)

	result, err := db.Exec("insert into client (firstname, lastname, street, city, statec, zip, email, username, pass) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)", firstname, lastname, street, city, state, zip, email, username, pass)
	if err != nil {
		log.Fatal(err)
	}
	temp, err := template.ParseFiles("templates/confirmation.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
	fmt.Println(result)
}

func login(w http.ResponseWriter, r *http.Request) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	usernameClient = r.FormValue("username")
	passClient = r.FormValue("pass")

	fmt.Println("Username Login", usernameClient)
	fmt.Println("password Login", passClient)
	username1 := functions.SearchByUsername(db, usernameClient)
	password1 := functions.SearchPassByUsername(db, usernameClient)
	fmt.Println(username1)
	fmt.Println(password1)
	if usernameClient == username1 {
		if passClient == password1 {
			temp, err := template.ParseFiles("templates/user.html")
			if err != nil {
				log.Fatal(err)
			}
			temp.Execute(w, nil)
			//fmt.Println("Username and password are correct")
		} else {
			temp, err := template.ParseFiles("templates/reject.html")
			if err != nil {
				log.Fatal(err)
			}
			temp.Execute(w, nil)
		}
	} else {
		temp, err := template.ParseFiles("templates/reject.html")
		if err != nil {
			log.Fatal(err)
		}
		temp.Execute(w, nil)
		//fmt.Println("Incorrect")
	}
}

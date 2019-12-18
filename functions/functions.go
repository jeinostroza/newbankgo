package functions

import (
	"database/sql"
	"fmt"
)

//SearchPassByUsername search by username and return password
func SearchPassByUsername(db *sql.DB, searchvalue string) string {
	row := db.QueryRow("select pass from client where username = $1", searchvalue)
	var pass string
	row.Scan(&pass)
	return pass

}

//SearchByUsername search by username and return password
func SearchByUsername(db *sql.DB, searchvalue string) string {
	row := db.QueryRow("select username from client where username = $1", searchvalue)
	var username string
	row.Scan(&username)
	return username

}

//GetAllClient pull information of a client
func GetAllClient(db *sql.DB, usernameget string) {
	row := db.QueryRow("select firstname, lastname, street, city, statec, zip, email, montlyincomes, monthlyexpenses from client where username = $1", usernameget)
	var firstname, lastname, street, city, statec, zip, email string
	var montlyincomes, monthlyexpenses float64
	row.Scan(&firstname, &lastname, &street, &city, &statec, &zip, &email, &montlyincomes, &monthlyexpenses)
	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(street)
	fmt.Println(city)
	fmt.Println(statec)
	fmt.Println(zip)
	fmt.Println(email)
	fmt.Println(montlyincomes)
	fmt.Println(monthlyexpenses)

}

//UpdateIncomes enter the income in client information
func UpdateIncomes(db *sql.DB, income float64, username string) {
	row := db.QueryRow(`update client set montlyincomes = $1 where username = $2`, income, username)
	var incomeup float64
	row.Scan(&incomeup)

}

//UpdateExpenses update the monthly expenses record
func UpdateExpenses(db *sql.DB, expenses float64, username string) {
	row := db.QueryRow(`update client set monthlyexpenses = $1 where username = $2`, expenses, username)
	var expensesup float64
	row.Scan(&expensesup)

}

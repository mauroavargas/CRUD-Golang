package main

import (
	// "log"
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (connect *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := "root"
	DBName := "go_system"

	connect, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1:8889)/"+DBName)
	if err != nil {
		panic(err.Error())
	}
	return connect
}

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/insert", Insert)
	fmt.Println("Server running!")
	http.ListenAndServe(":8080", nil)
}

type Employee struct {
	Id    int
	Name  string
	Email string
}

func Index(w http.ResponseWriter, r *http.Request) {

	establishedConnection := connectDB()
	getRecords, err := establishedConnection.Query("SELECT * FROM employees")

	if err != nil {
		panic(err.Error())
	}

	employee := Employee{}
	arrayEmployee := []Employee{}

	for getRecords.Next() {
		var id int
		var name, email string
		err = getRecords.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		employee.Id = id
		employee.Name = name
		employee.Email = email

		arrayEmployee = append(arrayEmployee, employee)

	}

	// fmt.Println(arrayEmployee)

	// fmt.Fprintf(w, "Hello World!")
	templates.ExecuteTemplate(w, "index", arrayEmployee)
}

func Add(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "add", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		name := r.FormValue("name")
		email := r.FormValue("email")

		establishedConnection := connectDB()
		addRecords, err := establishedConnection.Prepare("INSERT INTO employees(name, email) VALUES(?, ?)")

		if err != nil {
			panic(err.Error())
		}
		addRecords.Exec(name, email)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}
}

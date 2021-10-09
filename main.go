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

func Index(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Hello World!")
	templates.ExecuteTemplate(w, "index", nil)
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

		http.Redirect(w, r, "/", 301)

	}
}

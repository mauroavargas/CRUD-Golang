package main

import (
	// "log"
	"fmt"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Index)
	fmt.Println("Server running!")
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World!")
	templates.ExecuteTemplate(w, "index", nil)
}

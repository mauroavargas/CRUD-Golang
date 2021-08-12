package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Index)
	log.Println("Server running!")
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

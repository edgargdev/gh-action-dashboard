package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/partial.html",
))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func htmxHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Message": "Hello from HTMX!",
	}
	err := templates.ExecuteTemplate(w, "partial.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/htmx", htmxHandler)

	port := ":8080"
	log.Println("Server is running on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}


package handlers

import (
	"groupie-tracker/data"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Artists []data.Artist
}

var (
	homeTemplate *template.Template
)

func init() {
	var err error
	homeTemplate, err = template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatalf("Error parsing home template: %v", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request, artists []data.Artist) {
	if r.URL.Path != "/" {
		log.Printf("Page not found: %s", r.URL.Path)
		ErrorHandler(w, r, http.StatusNotFound, "The page you are looking for does not exist.")
		return
	}

	data := PageData{Artists: artists}
	if err := homeTemplate.Execute(w, data); err != nil {
		log.Printf("Error executing home template: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError, "An error occurred while processing your request.")
	}
}

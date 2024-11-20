package handlers

import (
	"groupie-tracker/data"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Loading bool
	Artists []data.Artist
}

// Global variable to store parsed template (parsed once at server startup)
var tmpl *template.Template

// Initialize the template once when the application starts
func init() {
	var err error
	tmpl, err = template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
}

// HomeHandler handles the home page rendering
func HomeHandler(w http.ResponseWriter, r *http.Request, artists []data.Artist) {
	err := tmpl.Execute(w, PageData{Loading: false, Artists: artists})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

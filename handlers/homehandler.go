package handlers

import (
	"groupie-tracker/data"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	State   data.State
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
	// Convert artists from data.Artist to Artist
	artistData := []Artist{}
	for _, artist := range artists {
		artistData = append(artistData, Artist{
			Name:      artist.Name,
			ImageLink: artist.Image, // Assuming Image field in data.Artist
		})
	}

	// Execute the template with the artist data
	err := tmpl.Execute(w, artistData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package handlers

import (
	"groupie-tracker/data"
	"html/template"
	"log"
	"net/http"
)

type ArtistPageData struct {
	Artist    data.Artist
	Locations data.Locations
	Relations data.Relations
}

func init() {
	var err error
	tmpl, err = template.ParseFiles("templates/artist.html")
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request, artist data.Artist, locations data.Locations, relations data.Relations) {
	data := ArtistPageData{Artist: artist, Locations: locations, Relations: relations}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

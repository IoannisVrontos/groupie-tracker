package handlers

import (
	"groupie-tracker/data"
	"log"
	"net/http"
)

type ArtistPageData struct {
	Artist            data.Artist
	LocationsAndDates []string
}

func ArtistHandler(w http.ResponseWriter, r *http.Request, artists []data.Artist, id int) {
	if id < 0 || id > len(artists) {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Build artist page data
	selectedArtist := artists[id-1]
	data := ArtistPageData{
		Artist: selectedArtist,
	}

	// Render artist template
	if err := artistTemplate.Execute(w, data); err != nil {
		log.Printf("Error executing artist template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

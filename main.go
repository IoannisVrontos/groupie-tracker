package main

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"net/http"
	"os"
)

func main() {
	var artists []data.Artist
	fetchedArtists, err := data.InitializeData()
	if err != nil {
		fmt.Println("Failed to fetch artists:", err)
	} else {
		artists = fetchedArtists
	}

	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(artists); err != nil {
			handlers.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to encode artists data.")
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, artists)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	fmt.Printf("Server is running on port %s...\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"net/http"
	"strconv"
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

	fmt.Printf("\rServer is running on port http://localhost:8080 ...\n")
	port := 8080
	err = http.ListenAndServe("localhost:8080", nil)
	for err != nil {
		port++
		fmt.Printf("\033[A\rServer is running on port http://localhost:%d ...\n", port)
		err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	}
}

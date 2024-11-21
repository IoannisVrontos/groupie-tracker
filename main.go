package main

import (
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"net/http"
	"strconv"
)

func main() {
	currentState := data.Loading
	var artists []data.Artist
	fetchedArtists, err := data.InitializeData()
	if err != nil {
		currentState = data.Error
		fmt.Println("Failed to fetch artists:", err)
	} else {
		artists = fetchedArtists
		currentState = data.Success
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, currentState, artists)
	})

	http.HandleFunc("/artist/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Path[len("/artist/"):])
		if err != nil {
			http.Error(w, "Invalid artist ID", http.StatusBadRequest)
			return
		}
		if id < 1 || id > len(artists) {
			http.Error(w, "Artist not found", http.StatusNotFound)
			return
		}
		handlers.ArtistHandler(w, r, artists, id)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Printf("\rServer is running on port http://localhost:8080 ...\n")
	port := 8080
	err = http.ListenAndServe(":8080", nil)
	for err != nil {
		port++
		fmt.Printf("\033[A\rServer is running on port http://localhost:%d ...\n", port)
		err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	}
}

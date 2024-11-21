package main

import (
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	currentState := data.Loading
	var artists []data.Artist
	var mu sync.RWMutex

	// Run GetArtists in a goroutine
	go func() {
		fetchedArtists, err := data.GetArtists()
		mu.Lock()
		if err != nil {
			currentState = data.Error
			fmt.Println("Failed to fetch artists:", err)
		} else {
			artists = fetchedArtists
			currentState = data.Success
		}
		mu.Unlock()
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		defer mu.RUnlock()
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

	http.ListenAndServe(":8080", nil)
}

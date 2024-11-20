package main

import (
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"net/http"
	"strconv"
)



func main() {

	artistsChan := make(chan []data.Artist)
	errChan := make(chan error)

	currentState := data.Loading
		var artists []data.Artist
		var err error

	go func() {
		artists, err = data.GetArtists()
		if err != nil {
			errChan <- err
			currentState = data.Error
			return
		}
		artistsChan <- artists
		currentState = data.Success

	}()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.HomeHandler(w, r,currentState, artists)
		})


	http.HandleFunc("/artist/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Path[len("/artist/"):])
		if err != nil {
			http.Error(w, "Invalid artist ID", http.StatusBadRequest)
			return
		}
		handlers.ArtistHandler(w, r, artists[id-1], artists[id-1].Locations, artists[id-1].Relations)
	})

	http.ListenAndServe(":8080", nil)
}

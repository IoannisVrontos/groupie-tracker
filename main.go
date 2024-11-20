package main

import (
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"net/http"
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


	http.ListenAndServe(":8080", nil)
}

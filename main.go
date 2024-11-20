package main

import (
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"net/http"
	"strconv"
)

func main() {
	artists, err := data.GetArtists()
	if err != nil {
		fmt.Println("Error getting artists: ", err)
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, artists)
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

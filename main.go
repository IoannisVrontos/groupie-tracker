package main

import (
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"net/http"
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

	http.ListenAndServe(":8080", nil)
}
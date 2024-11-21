package handlers

import (
	"groupie-tracker/data"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	State   data.State
	Artists []data.Artist
}

var (
	homeTemplate   *template.Template
	artistTemplate *template.Template
)

func init() {
	var err error
	homeTemplate, err = template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatalf("Error parsing home template: %v", err)
	}

	artistTemplate, err = template.ParseFiles("templates/artist.html")
	if err != nil {
		log.Fatalf("Error parsing artist template: %v", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request, state data.State, artists []data.Artist) {
	data := PageData{State: state, Artists: artists}

	if err := homeTemplate.Execute(w, data); err != nil {
		log.Printf("Error executing home template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

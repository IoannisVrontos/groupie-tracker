package data

import (
	"fmt"
)

type Artist struct {
	ID               int      `json:"id"`
	Image            string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Locations        Locations
	ConcertDates     Dates
	Relations        Relations
	LocationsLink    string `json:"locations"`
	ConcertDatesLink string `json:"concertDates"`
	RelationsLink    string `json:"relations"`
}

func (a *Artist) Init() {
	locations, err := GetLocations(a.LocationsLink)
	if err != nil {
		fmt.Println("Error getting locations: ", err)
	}
	a.Locations = locations
	dates, err := GetDates(a.ConcertDatesLink)
	if err != nil {
		fmt.Println("Error getting dates: ", err)
	}
	a.ConcertDates = dates
	relations, err := GetRelations(a.RelationsLink)
	if err != nil {
		fmt.Println("Error getting relations: ", err)
	}
	a.Relations = relations
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesLink string   `json:"dates"`
	Dates     Dates
}

func (l *Locations) Init() {
	dates, err := GetDates(l.DatesLink)
	if err != nil {
		fmt.Println("Error getting dates: ", err)
	}
	l.Dates = dates
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type State int

const (
    Loading State = iota // Loading == 0
    Success              // Success == 1
    Error                // Error == 2
)

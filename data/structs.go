package data

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetLocations(link string) (Locations, error) {
	var locations Locations
	resp, err := http.Get(link)
	if err != nil {
		return Locations{}, fmt.Errorf("failed to get locations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
			return Locations{}, fmt.Errorf("failed to decode locations: %w", err)
		}
	}
	return locations, nil
}

func GetDates(link string) (Dates, error) {
	var dates Dates
	resp, err := http.Get(link)
	if err != nil {
		return Dates{}, fmt.Errorf("failed to get dates: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&dates); err != nil {
		return Dates{}, fmt.Errorf("failed to decode dates: %w", err)
	}

	return dates, nil
}

func GetRelations(link string) (Relations, error) {
	var relations Relations
	resp, err := http.Get(link)
	if err != nil {
		return Relations{}, fmt.Errorf("failed to get relations: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return Relations{}, fmt.Errorf("failed to decode relations: %w", err)
	}

	return relations, nil
}

func GetArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return []Artist{}, fmt.Errorf("failed to get artists: %w", err)
	}
	defer resp.Body.Close()

	var artists []Artist

	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return []Artist{}, fmt.Errorf("failed to decode artists: %w", err)
	}

	for i := range artists {
		artists[i].Init()
	}

	return artists, nil
}

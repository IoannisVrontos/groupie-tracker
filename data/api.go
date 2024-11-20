package data

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
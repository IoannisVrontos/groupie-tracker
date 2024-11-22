package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetDates() (AllDates, error) {
	var dates AllDates
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return AllDates{}, fmt.Errorf("failed to get dates: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&dates); err != nil {
		return AllDates{}, fmt.Errorf("failed to decode dates: %w", err)
	}

	return dates, nil
}
func GetLocations() (AllLocations, error) {
	var locations AllLocations
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return AllLocations{}, fmt.Errorf("failed to get locations: %w", err)
	}
	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&locations)
		if err != nil {
			return AllLocations{}, fmt.Errorf("failed to decode locations: %w", err)
		}
	}
	defer resp.Body.Close()
	return locations, nil
}

func GetRelations() (AllRelations, error) {
	var relations AllRelations
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return AllRelations{}, fmt.Errorf("failed to get relations: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return AllRelations{}, fmt.Errorf("failed to decode relations: %w", err)
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

	return artists, nil
}

func InitializeData() ([]Artist, error) {
	fetchedArtists, err := GetArtists()
	if err != nil {
		return []Artist{}, err
	}
	relations, err := GetRelations()
	if err != nil {
		return []Artist{}, err
	}
	for i := range relations.Index {
		relations.Index[i].DatesLocations = mapDatesLocations(relations.Index[i].DatesLocations)
	}
	for i := range fetchedArtists {
		fetchedArtists[i].Relations = relations.Index[i].DatesLocations
	}
	return fetchedArtists, nil
}

func mapDatesLocations(datesLocations map[string][]string) map[string][]string {
	newDatesLocations := make(map[string][]string)
	for k, v := range datesLocations {
		newKey := strings.ReplaceAll(k, "_", " ")
		newKey = strings.ReplaceAll(newKey, "-", ", ")
		newKey = strings.ReplaceAll(newKey, newKey, capitalizeWords(newKey))
		newKey = strings.ReplaceAll(newKey, "Uk", "UK")
		newKey = strings.ReplaceAll(newKey, "Usa", "USA")
		newDatesLocations[newKey] = v
	}
	return newDatesLocations
}

func capitalizeWords(sentence string) string {
	words := strings.Fields(sentence)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	return strings.Join(words, " ")
}

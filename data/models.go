package data

type Artist struct {
	ID               int      `json:"id"`
	Image            string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Relations        map[string][]string
	LocationsLink    string `json:"locations"`
	ConcertDatesLink string `json:"concertDates"`
	RelationsLink    string `json:"relations"`
}

type AllLocations struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesLink string   `json:"dates"`
	Dates     Dates
}

type AllRelations struct {
	Index []Relations `json:"index"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type AllDates struct {
	Index []Dates `json:"index"`
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

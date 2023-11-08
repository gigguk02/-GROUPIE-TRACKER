package art

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int32    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
type Relation struct {
	Id              int                 `json:"id"`
	LocationAndDate map[string][]string `json:"datesLocations"`
}
type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
	DatesUrl string   `json:"dates"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

func AllArtists() ([]Artist, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("6")
		return []Artist{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("7")
		return []Artist{}, http.StatusInternalServerError
	}
	var Artists []Artist
	err1 := json.Unmarshal(body, &Artists)
	if err1 != nil {
		fmt.Println("8")
		return []Artist{}, http.StatusInternalServerError
	}
	return Artists, http.StatusOK

}

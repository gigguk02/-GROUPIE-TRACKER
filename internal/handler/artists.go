package handler

import (
	"artists/models"
	"artists/pkg"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func AllArtists() ([]models.Artist, int, int, int, int, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return []models.Artist{}, 0, 0, 0, 0, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []models.Artist{}, 0, 0, 0, 0, http.StatusInternalServerError
	}

	err3 := json.Unmarshal(body, &Artists)
	// fmt.Println(Artists)
	if err3 != nil {
		return []models.Artist{}, 0, 0, 0, 0, http.StatusInternalServerError
	}
	relation, statusCode := Relations()
	if statusCode != http.StatusOK {
		return []models.Artist{}, 0, 0, 0, 0, http.StatusInternalServerError
	}
	for i := range Artists {
		Artists[i].Relation = relation.Index[i].LocationAndDate
	}

	for i := range Artists {
		newLocationAndDate := make(map[string][]string)
		for key, value := range Artists[i].Relation {
			newKey := strings.Title(strings.ReplaceAll(key, "_", " "))
			newLocationAndDate[newKey] = value

		}
		Artists[i].Relation = newLocationAndDate
	}
	errStatus, maxCreateDate, minCreateDate, maxFirstAlbum, minFirstAlbum := pkg.MaxMin(Artists)
	if errStatus != http.StatusOK {
		return []models.Artist{}, 0, 0, 0, 0, http.StatusBadRequest
	}

	return Artists, maxCreateDate, minCreateDate, maxFirstAlbum, minFirstAlbum, http.StatusOK
}

func Relations() (models.Relations, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil || res.StatusCode != http.StatusOK {
		return models.Relations{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Relations{}, http.StatusInternalServerError
	}
	var Relations models.Relations
	json.Unmarshal(body, &Relations)

	return Relations, http.StatusOK
}

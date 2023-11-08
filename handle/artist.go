package handle

import (
	"artists/art"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil || id > 52 || id < 1 {
		fmt.Println(0)
		Errors(w, http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodGet {
		fmt.Println(1)
		Errors(w, http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("./html/artist.html")
	if err != nil {
		fmt.Println(2)
		Errors(w, http.StatusInternalServerError)
		return
	}
	Artist, statusCode := OneArtist(idstr)
	if statusCode != http.StatusOK {
		fmt.Println(3)
		Errors(w, statusCode)
		return
	}
	Relation, statusCode := Relation(idstr)
	if statusCode != http.StatusOK {
		fmt.Println(4)
		Errors(w, statusCode)
		return
	}
	Location, statusCode := Location(idstr)
	if statusCode != http.StatusOK {
		fmt.Println(5)
		Errors(w, statusCode)
		return
	}
	Dates, statusCode := Dates(idstr)
	if statusCode != http.StatusOK {
		fmt.Println(6)
		Errors(w, statusCode)
		return
	}
	// newLocationAndDate := make(map[string][]string)

	newLocationAndDate := make(map[string][]string)
	for key, value := range Relation.LocationAndDate {
		newKey := strings.Title(strings.ReplaceAll(key, "_", " "))
		newLocationAndDate[newKey] = value

	}
	Relation.LocationAndDate = newLocationAndDate

	inf := struct {
		Art  art.Artist
		Rel  art.Relation
		Loc  art.Location
		Date art.Date
	}{
		Art:  Artist,
		Rel:  Relation,
		Loc:  Location,
		Date: Dates,
	}
	if err := tmp.Execute(w, inf); err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}

}
func OneArtist(idstr string) (art.Artist, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + idstr)
	if err != nil || res.StatusCode != http.StatusOK {
		return art.Artist{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return art.Artist{}, http.StatusInternalServerError
	}
	var OneArtist art.Artist
	json.Unmarshal(body, &OneArtist)
	return OneArtist, http.StatusOK

}
func Relation(idstr string) (art.Relation, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + idstr)
	if err != nil || res.StatusCode != http.StatusOK {
		return art.Relation{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return art.Relation{}, http.StatusInternalServerError
	}
	var Relation art.Relation
	json.Unmarshal(body, &Relation)
	return Relation, http.StatusOK

}
func Location(idstr string) (art.Location, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + idstr)
	if err != nil || res.StatusCode != http.StatusOK {
		return art.Location{}, http.StatusInternalServerError

	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return art.Location{}, http.StatusInternalServerError
	}
	var Locations art.Location
	json.Unmarshal(body, &Locations)
	return Locations, http.StatusOK

}
func Dates(idstr string) (art.Date, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + idstr)
	if err != nil || res.StatusCode != http.StatusOK {
		return art.Date{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return art.Date{}, http.StatusInternalServerError
	}
	var Dates art.Date
	json.Unmarshal(body, &Dates)
	return Dates, http.StatusOK

}

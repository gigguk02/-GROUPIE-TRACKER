package handler

import (
	"artists/models"
	"artists/pkg"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var (
	Artists          []models.Artist
	WithoutDouble    []string
	NumWithoutDouble []int
	MinCreateDate    int
	MaxCreateDate    int
	MinFirstAlbum    int
	MaxFirstAlbum    int
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound)
		// http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errors(w, http.StatusMethodNotAllowed)
		// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("./ui/template/main.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var errStatus int
	Artists, MinCreateDate, MaxCreateDate, MaxFirstAlbum, MinFirstAlbum, errStatus = AllArtists()
	WithoutDoubleDateLoc()

	if errStatus != http.StatusOK {
		Errors(w, errStatus)
		// http.Error(w, http.StatusText(errStatus), errStatus)
		return

	}
	if err := tmp.Execute(w, models.Art{
		Art:           Artists,
		SearchArtist:  Artists,
		Relation:      WithoutDouble,
		Client:        "",
		CreationDate:  NumWithoutDouble,
		MinCreateDate: MinCreateDate,
		MaxCreateDate: MaxCreateDate,
		MinFirstAlbum: MinFirstAlbum,
		MaxFirstAlbum: MaxFirstAlbum,
	}); err != nil {
		fmt.Println("5")
		Errors(w, http.StatusInternalServerError)
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil || id > 52 || id < 1 {
		Errors(w, http.StatusBadRequest)
		// http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodGet {
		Errors(w, http.StatusMethodNotAllowed)
		// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("./ui/template/artist.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := tmp.Execute(w, Artists[id-1]); err != nil {
		Errors(w, http.StatusInternalServerError)
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func SearchPages(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		Errors(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		Errors(w, http.StatusMethodNotAllowed)
		return
	}
	clientWords := r.FormValue("search")
	SearchArtists, StatusCode := pkg.Search(clientWords, Artists)
	if StatusCode != http.StatusOK {
		Errors(w, StatusCode)
		return
	}
	tmp, err := template.ParseFiles("./ui/template/main.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}

	if err := tmp.Execute(w, models.Art{
		Art:           Artists,
		SearchArtist:  SearchArtists,
		Relation:      WithoutDouble,
		Client:        "",
		CreationDate:  NumWithoutDouble,
		MinCreateDate: MinCreateDate,
		MaxCreateDate: MaxCreateDate,
		MinFirstAlbum: MinFirstAlbum,
		MaxFirstAlbum: MaxFirstAlbum,
	}); err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}
}
func Filter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter" {
		Errors(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		Errors(w, http.StatusMethodNotAllowed)
		return
	}
	createDateStart := r.FormValue("createDateStart")
	createDateEnd := r.FormValue("createDateEnd")
	firstAlbumStart := r.FormValue("firstAlbumStart")
	firstAlbumEnd := r.FormValue("firstAlbumEnd")
	location := r.FormValue("location")
	numberOfMembers := r.Form["member"]
	tmp, err := template.ParseFiles("./ui/template/main.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}

	// fmt.Println("Create Date Start:", createDateStart)
	// fmt.Println("Create Date End:", createDateEnd)
	// fmt.Println("First Album Start:", firstAlbumStart)
	// fmt.Println("First Album End:", firstAlbumEnd)
	// fmt.Println("Location:", location)
	// fmt.Println("Number:", numberOfMembers)

	artists, err1 := pkg.Filter(createDateStart, createDateEnd, firstAlbumStart, firstAlbumEnd, location, numberOfMembers, Artists)
	if err1 != http.StatusOK {
		Errors(w, err1)
		return

	}
	if err := tmp.Execute(w, models.Art{
		Art:           Artists,
		SearchArtist:  artists,
		Relation:      WithoutDouble,
		Client:        "",
		CreationDate:  NumWithoutDouble,
		MinCreateDate: MinCreateDate,
		MaxCreateDate: MaxCreateDate,
		MinFirstAlbum: MinFirstAlbum,
		MaxFirstAlbum: MaxFirstAlbum,
	}); err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}

}

func WithoutDoubleDateLoc() {
	var (
		newLocation []string
		allDate     []int
	)
	for i := range Artists {
		for key := range Artists[i].Relation {
			newLocation = append(newLocation, key)
		}
		allDate = append(allDate, int(Artists[i].CreationDate))
	}
	for i := range allDate {
		found := false
		for j := range NumWithoutDouble {
			if allDate[i] == NumWithoutDouble[j] {
				found = true
				break
			}
		}
		if !found {
			NumWithoutDouble = append(NumWithoutDouble, allDate[i])
		}
	}
	for i := range newLocation {
		found := false
		for j := range WithoutDouble {
			if newLocation[i] == WithoutDouble[j] {
				found = true
				break
			}
		}
		if !found {
			WithoutDouble = append(WithoutDouble, newLocation[i])

		}
	}

}

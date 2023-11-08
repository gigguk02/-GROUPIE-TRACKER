package handle

import (
	"artists/art"
	"fmt"
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Errors(w, http.StatusInternalServerError)
		fmt.Println("1")
		return
	}
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound)
		return
	}
	tmp, err := template.ParseFiles("./html/main.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	artists, errStatus := art.AllArtists()
	if errStatus != http.StatusOK {
		Errors(w, http.StatusInternalServerError)
		fmt.Println("4")
		return
	}
	if err := tmp.Execute(w, artists); err != nil {
		Errors(w, http.StatusInternalServerError)
		fmt.Println("5")
		return
	}

}

package handler

import (
	"log"
	"net/http"
)

func Handler() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", MainPage)
	mux.HandleFunc("/artist", ArtistPage)
	mux.HandleFunc("/search", SearchPages)
	mux.HandleFunc("/filter", Filter)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	log.Println("Запуск веб-сервера на http://localhost:8080")
	return http.ListenAndServe(":8080", mux)

}

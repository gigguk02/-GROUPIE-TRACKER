package handle

import (
	"net/http"
	"text/template"
)

func Errors(w http.ResponseWriter, code int) {
	errors := struct {
		ErrorCode int
		ErrorText string
	}{
		ErrorCode: code,
		ErrorText: http.StatusText(code),
	}
	tmpl, err := template.ParseFiles("html/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}
	if err1 := tmpl.Execute(w, errors); err1 != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}

}

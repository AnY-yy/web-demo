package controller

import (
	"net/http"
	"text/template"
)

func registerIndexRoutes() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("page/index.html")

	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

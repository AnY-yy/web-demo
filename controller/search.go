package controller

import (
	"net/http"
	"text/template"
)

func registerSearchRoutes() {
	http.HandleFunc("/search", handleSearch)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("page/search_student.html")

	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

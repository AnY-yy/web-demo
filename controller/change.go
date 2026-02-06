package controller

import (
	"net/http"
	"text/template"
)

func registerChangeRoutes() {
	http.HandleFunc("/change", handleChange)
}

func handleChange(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("page/change_student.html")

	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

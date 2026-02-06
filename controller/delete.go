package controller

import (
	"net/http"
	"text/template"
)

func registerDelete() {
	http.HandleFunc("/delete", handleDelete)
}

func handleDelete(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("page/delete_student.html")

	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

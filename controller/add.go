package controller

import (
	"net/http"
	"text/template"
)

func registerAddRoutes() {
	http.HandleFunc("/add", handleAdd)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	t, _ := template.ParseFiles("page/add_student.html")
	// 返回到ResponseWriter中
	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

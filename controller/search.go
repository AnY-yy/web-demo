package controller

import (
	"Minimalist_Web_Application--StudentInfo_Managment_System/db"
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

	if r.Method == http.MethodPost {
		// 解析表单
		err = r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 获得表单数据
		idOrName := r.Form.Get("IdOrName")
		info := &db.Info{}
		err = info.SearchStudentInfo(idOrName)
		if err != nil {
			showAlert(w, "Search data failed!"+err.Error(), "/search")
		} else {
			str := "id:" + info.Id + "Name:" + info.Name + "Gender:" + info.Gender + "Age:" + info.Age + "Phone:" + info.Phone + "Address:" + info.Address
			showAlert(w, "Search data succeed!"+str, "/search")
		}
	}
}

package controller

import (
	"Minimalist_Web_Application--StudentInfo_Managment_System/db"
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

	if r.Method == http.MethodPost {
		// 解析表单
		err = r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 获取表单数据
		info := db.Info{
			Id:      r.Form.Get("id"),
			Name:    r.Form.Get("name"),
			Gender:  r.Form.Get("gender"),
			Age:     r.Form.Get("age"),
			Phone:   r.Form.Get("phone"),
			Address: r.Form.Get("address"),
		}
		err := info.UpdateStudentInfo()
		if err != nil {
			showAlert(w, "Update data failed!"+err.Error(), "/change")
		} else {
			showAlert(w, "Update data success!", "/change")
		}
	}
}

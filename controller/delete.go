package controller

import (
	"Minimalist_Web_Application--StudentInfo_Managment_System/db"
	"net/http"
	"text/template"
)

func registerDelete() {
	http.HandleFunc("/delete", handleDelete)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("page/delete_student.html")

	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		err = db.DeleteStudentInfo(idOrName)
		if err != nil {
			showAlert(w, "Delete data failed!"+err.Error(), "/delete")
		} else {
			showAlert(w, "Delete student success!", "/delete")
		}
	}
}

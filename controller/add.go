package controller

import (
	"Minimalist_Web_Application--StudentInfo_Managment_System/db"
	"net/http"
	"text/template"
)

func registerAddRoutes() {
	http.HandleFunc("/add", handleAdd)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	t, _ := template.ParseFiles("page/add_student.html")
	// 渲染模板文件 并返回给请求中
	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 仅点击确定按钮才获取表单数据 即POST请求
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
		// 插入到数据库中
		err = info.InsertStudentInfo()
		if err != nil {
			showAlert(w, "insert data failed!"+err.Error(), "/add")
		} else {

			showAlert(w, "Search the data!", "/add")
		}
	}
}

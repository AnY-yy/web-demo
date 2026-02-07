package main

import (
	"Minimalist_Web_Application--StudentInfo_Managment_System/controller"
	"Minimalist_Web_Application--StudentInfo_Managment_System/db"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// 启动数据库连接
	db.Connect()

	// 路由注册
	controller.RegisterRoutes()

	// 监听8000端口 使用goroutine运行 不影响其他中间件
	go http.ListenAndServe(":8000", nil)

	// 监听服务器端口
	err := http.ListenAndServeTLS("localhost:8080", "certificate/cert.pem", "certificate/key.pem", nil)
	if err != nil {
		panic(err)
	}
}

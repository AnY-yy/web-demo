package db

import "fmt"

// 判断是否连接成功
func judgeConnect(err error) {
	if err != nil {
		fmt.Println("Successfully connected!", err)
		panic(err)
	} else {
		fmt.Println("Successfully connected!")
		return
	}
}

package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Connect() {
	// 连接数据库
	dbStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)

	db, err = sql.Open("mysql", dbStr)
	if err != nil {
		panic(err)
	}
	// 使用生成的上下文变量判断是否连接成功
	ctx := context.Background()
	judgeConnect(db.PingContext(ctx))
}

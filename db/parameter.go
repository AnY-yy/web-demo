package db

// 连接数据库用的参数
const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "aaa2021s"
	dbName   = "WebTest"
)

// Info ..
type Info struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Age     string `json:"age"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

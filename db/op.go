package db

import (
	"errors"
	"unicode"
)

func (i *Info) InsertStudentInfo() error {
	// 先判断是否存在这个id
	err := i.SearchStudentInfo(i.Id)
	if err == nil {
		err = errors.New("the id is already in use")
		return err
	}
	// 存放sql语句字符串
	sqlStr := "INSERT INTO student (id, name, gender, age, phone, address) VALUES(?, ?, ?, ?, ?, ?)"
	// 使用Prepare传入参数
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(i.Id, i.Name, i.Gender, i.Age, i.Phone, i.Address)

	return err
}

func DeleteStudentInfo(s string) error {
	sqlStr := "DELETE FROM student WHERE "
	if isDigits(s) {
		sqlStr = sqlStr + "id=" + s
	} else {
		sqlStr = sqlStr + "name=" + s
	}
	_, err := db.Exec(sqlStr)

	return err
}

func (i *Info) UpdateStudentInfo() error {
	// 判断是否存在该信息
	sqlStr := "SELECT name FROM student WHERE id=?"
	var name string
	err := db.QueryRow(sqlStr, i.Id).Scan(&name)
	if err != nil {
		return err
	}
	if len(name) < 0 {
		err = errors.New("the data is empty")
		return err
	}
	// 更新全部信息
	sqlStr = "UPDATE student SET name=?, gender=?, age=?, phone=?, address=? WHERE id=?"
	_, err = db.Exec(sqlStr, i.Name, i.Gender, i.Age, i.Phone, i.Address, i.Id)

	return err
}

func (i *Info) SearchStudentInfo(s string) error {
	sqlStr := "SELECT id, name, gender, age, phone, address from student where "
	if isDigits(s) {
		sqlStr = sqlStr + "id=" + s
	} else {
		sqlStr = sqlStr + "name=" + s
	}
	err := db.QueryRow(sqlStr).Scan(&i.Id, i.Name, &i.Gender, &i.Age, &i.Phone, &i.Address)

	return err
}

// 判断字符串是不是由纯数字组成 学号是纯数字 名字不是
func isDigits(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 遍历字符串
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

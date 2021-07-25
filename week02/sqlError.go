package week02

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var Db *sql.DB

func init() {
	var err error
	dsn := "root:123456@tcp(192.168.6.136:3306)/test?charset=utf8mb4&parseTime=True"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(errors.Wrap(err, "数据库连接失败"))
	}
}

func QuertById(id int) (*User, error) {
	var user User
	sqlStr := "select id, name, sex, age from user where id=?"
	row := Db.QueryRow(sqlStr, id)
	err := row.Scan(&user.Id, &user.Name, &user.Sex, &user.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "查询失败")
	}
	return &user, nil
}

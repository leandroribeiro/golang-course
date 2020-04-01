package repository

import "github.com/jmoiron/sqlx"
import _ "github.com/go-sql-driver/mysql"

var Db *sqlx.DB

func Open() (err error){
	err = nil
	Db, err = sqlx.Open("mysql", "root:MySql2019!@tcp(127.0.0.1:3306)/GoCourse")
	if err != nil {
		return
	}
	err = Db.Ping()
	return
}
package models

import (
    "fmt"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Initer _Init

type _Init struct {
	_db_url string
	_db *sql.DB
}


func (_init *_Init) Init() int {
	_init._db_url = "guest:123456@(liudonghua.net:3306)/kone_guest?charset=utf8"
	var err error
	_init._db, err = sql.Open("mysql", _init._db_url)
	if(err != nil){
		fmt.Println("mysql open failed, error: ", err)
		return -1
	}
	return 0
}

package models

import (
    "fmt"
    "strconv"
    //"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Guest MGuest

type MGuest struct {

	Table_name string

	Guest_id int
	Name string
	Create_time string
	Visit_date string
	Face_id string
	Visit_building string
	Visit_floor string
	Visit_room string
}

func (guest *MGuest) Init() {
	guest.Table_name = "guest"
}

func (guest *MGuest) Save(guest_info *MGuest) int64 {

	fmt.Println("save guest info")

	prepare_sql := "INSERT into " + guest.Table_name + " SET " +
					" name=?" +
					",create_time=?" +
					",visit_date=?" +
					",face_id=?" +
					",visit_building=?" +
					",visit_floor=?" +
					",visit_room=?"

	stmt, err := Initer._db.Prepare(prepare_sql)

	if(err != nil){
		fmt.Println("sql prepare failed, error: ", err)
		return -1
	}

	res, err := stmt.Exec(
					guest_info.Name,
					guest_info.Create_time,
					guest_info.Visit_date,
					guest_info.Face_id,
					guest_info.Visit_building,
					guest_info.Visit_floor,
					guest_info.Visit_room )

	if(err != nil){
		fmt.Println("sql exec failed, error: ", err)
		return -1
	}

	guest_id, err := res.LastInsertId()

	if(err != nil){
		fmt.Println("get LastInsertId failed, error: ", err)
		return -1
	}

	return guest_id
}

func (guest *MGuest) Get(id int) []MGuest {

	var guests []MGuest;
	var guest_info MGuest;

	prepare_sql := "SELECT * from " + guest.Table_name + " where guest_id = " + strconv.Itoa(id)
	fmt.Println("sql: ", prepare_sql)

	/*
	stmt, err := model.Init._db.Prepare(prepare_sql)

	if(err != nil){
		fmt.Println("sql prepare failed, error: ", err)
		return -1
	}

	res, err := stmt.Exec(id)
	*/

	rows, err := Initer._db.Query(prepare_sql)

	if(err != nil){
		fmt.Println("sql query failed, error: ", err)
		return guests
	}


	for rows.Next() {

		err = rows.Scan(
					&guest_info.Guest_id,
					&guest_info.Name,
					&guest_info.Create_time,
					&guest_info.Visit_date,
					&guest_info.Face_id,
					&guest_info.Visit_building,
					&guest_info.Visit_floor,
					&guest_info.Visit_room )

		if(err != nil){
			fmt.Println("sql scan failed, error: ", err)
			return guests
		}

		guests = append(guests, guest_info)
	}

	return guests
}


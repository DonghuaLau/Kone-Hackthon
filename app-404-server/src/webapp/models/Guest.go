package models

import (
    "fmt"
    "strconv"
	_ "github.com/go-sql-driver/mysql"
)

var Guest MGuest
var VisitGuest VGuest

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

type VGuest struct{
	Guest_id int
	Is_visiting int
	Type  int
	Last_visit_time string
	//Position_info string
	Name string
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

func (guest *VGuest) GetVisitingGuests(_type int) []VGuest {
	var visiting_guests []VGuest;
	var visiting_guest_info VGuest;

	prepare_sql := "SELECT * from  visiting_guest where type = " + strconv.Itoa(_type) + " ORDER BY last_visit_time DESC"
	fmt.Println("sql: ", prepare_sql)


	rows, err := Initer._db.Query(prepare_sql)

	if(err != nil){
		fmt.Println("sql query failed, error: ", err)
		return visiting_guests
	}


	for rows.Next() {

		err = rows.Scan(
					&visiting_guest_info.Guest_id,
					&visiting_guest_info.Is_visiting,
					&visiting_guest_info.Last_visit_time,
					//&visiting_guest_info.Position_info,
					&visiting_guest_info.Type,
					&visiting_guest_info.Name)

		if(err != nil){
			fmt.Println("sql scan failed, error: ", err)
			return visiting_guests
		}

		visiting_guests = append(visiting_guests, visiting_guest_info)
	}

	return visiting_guests
}

/*
func (guest *VGuest) GetStrangers() []VGuest {
	var strangers []VGuest;
	var stranger_info VGuest;

	prepare_sql := "SELECT * from  visiting_guest where type=3 ORDER BY last_visit_time DESC"
	//fmt.Println("sql: ", prepare_sql)


	rows, err := Initer._db.Query(prepare_sql)

	if(err != nil){
		fmt.Println("sql query failed, error: ", err)
		return strangers
	}


	for rows.Next() {

		err = rows.Scan(
					&stranger_info.Guest_id,
					&stranger_info.Is_visiting,
					&stranger_info.Type,
					&stranger_info.Last_visit_time,
					&stranger_info.Name,
					&stranger_info.Position_info)

		if(err != nil){
			fmt.Println("sql scan failed, error: ", err)
			return strangers
		}

		strangers = append(strangers, stranger_info)
	}

	return strangers
}
*/

func (guest *VGuest)InsertVisitingGuest(faceid string, _type int) bool {

	query_sql := "SELECT guest_id from guest where face_id = '"+ faceid + "'"

	rows, err := Initer._db.Query(query_sql)

	if(err != nil){
		fmt.Println("sql query failed, error: ", err)
		return false
	}

	var guest_id int
	rows.Scan(&guest_id)

	insert_sql := "INSERT into visiting_guest values("+ strconv.Itoa(guest_id) + " , 1, "+ strconv.Itoa(_type) +", now())"

	rows,err = Initer._db.Query(insert_sql)

	if(err != nil) {
		fmt.Println("sql insert failed, error: ", err)
		return false
	}

	return true
}

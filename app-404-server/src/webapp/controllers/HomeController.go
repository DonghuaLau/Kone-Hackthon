package controllers

import (
	"fmt"
	"net/http"
	"webapp/models"
	"webapp/views"
)

var HomeController _HomeController

type _HomeController struct{}

func (ctl *_HomeController) Index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "HomeController index\n")
}

func (ctl *_HomeController) Profile(resp http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(resp, "HomeController profile\n")
}

func (ctl *_HomeController) CreateGuest(resp http.ResponseWriter, req *http.Request) {

	guest := models.MGuest{"-", 0,
				"Tom",
				"2017-10-15 08:30:00",
				"2017-10-16",
				"face_id_12135",
				"Kone Main Building",
				"3",
				"305" }
	models.Guest.Save(&guest)

	/*
	var guest MGuest
	guest._name = "Tom"
	guest._create_time = "2017-10-15 08:30:00"
	guest._visit_date = "2017-10-16"
	guest._face_id = "face_id_12135"
	guest._visit_building = "Kone Main Building"
	guest._visit_floor = "3"
	guest._visit_room = "305"
	models.Guest.save(guest)
	*/
}

func (ctl *_HomeController) GetGuests(resp http.ResponseWriter, req *http.Request) {


	/*
	guests := models.Guest.Get(1)
	size := len(guests)
	for i := 0; i < size; i++ {
		fmt.Println(guests[i])
	}
	*/

	staffs := models.VisitGuest.GetVisitingGuests(1)
	guests := models.VisitGuest.GetVisitingGuests(2)
	strangers := models.VisitGuest.GetVisitingGuests(3)


	views.HomeView.Show(resp, staffs, guests, strangers)
}

func (ctl *_HomeController) GetVisiting(resp http.ResponseWriter, req *http.Request) {


	guests := models.Guest.Get(1)
	size := len(guests)
	for i := 0; i < size; i++ {
		fmt.Println(guests[i])
	}

	//views.HomeView.Show(resp, guests)
}

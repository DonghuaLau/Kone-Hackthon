package views

import (
	"fmt"
	"net/http"
	"webapp/models"
	"html/template"
)

var HomeView _HomeView

type Info struct {
	Staffs []models.VGuest
	Guests []models.VGuest
	Strangers []models.VGuest
}

type _HomeView struct {
	Tmpl *template.Template
}

func (view *_HomeView) Init() {
	var err error
	view.Tmpl, err = template.ParseFiles("views/page.tmpl")
	if view.Tmpl == nil {
		fmt.Println("Template create failed, error: ", err)
	}
}

func (view *_HomeView) Show(resp http.ResponseWriter, staffs []models.VGuest, guests []models.VGuest, strangers []models.VGuest) {
	info := Info{staffs, guests, strangers}
	view.Tmpl.Execute(resp, info)
}

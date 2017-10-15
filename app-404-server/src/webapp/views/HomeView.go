package views

import (
	"fmt"
	"net/http"
	"webapp/models"
	"html/template"
)

var HomeView _HomeView

type Info struct {
	Guests []models.MGuest
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

func (view *_HomeView) Show(resp http.ResponseWriter, guests []models.MGuest) {
	info := Info{guests}
	view.Tmpl.Execute(resp, info)
}

package main

import (
	//"fmt"
	"log"
	"net/http"
	"webapp/controllers"
	"webapp/models"
	"webapp/views"
)



type WebApp struct {
	//_server  http.ServeMux
	_routers []Router
}

func (app *WebApp) init() {

	app.routers()


	// models init
	models.Initer.Init();
	models.Guest.Init();
	//models.VisitingGuest.init();
	//models.Staff.init();

	views.HomeView.Init();

	//for key, value := range app._routers {
	for _, value := range app._routers {
		//fmt.Println("key: ", key, ", method: ", value._method, ", URI: ", value._uri)
		http.HandleFunc(value._uri, value._handler)
	}
}

func (app *WebApp) start() {

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *WebApp) routers() {
	app._routers = []Router{
		{"GET", "/index",		controllers.HomeController.Index},
		{"GET", "/profile",		controllers.HomeController.Profile},
		{"GET", "/user/index",	controllers.UserController.Index},
		{"GET", "/user/profile",controllers.UserController.Profile},
		{"GET", "/guest/new",	controllers.HomeController.CreateGuest},
		{"GET", "/guests",		controllers.HomeController.GetGuests},
		{"GET", "/vistiting/list",		controllers.HomeController.GetVisiting},
	}
}

func main() {
	//test1()

	app := WebApp{}
	app.init()
	app.start()
}

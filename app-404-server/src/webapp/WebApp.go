package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/controllers"
)

type WebApp struct {
	//_server  http.ServeMux
	_routers []Router
}

func (app *WebApp) init() {

	app.routers()

	for key, value := range app._routers {
		fmt.Println("key: ", key, ", method: ", value._method, ", URI: ", value._uri)
		http.HandleFunc(value._uri, value._handler)
	}
}

func (app *WebApp) start() {

	err := http.ListenAndServe(":18080", nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *WebApp) routers() {
	app._routers = []Router{
		{"GET", "/index", controllers.HomeController.Index},
		{"GET", "/profile", controllers.HomeController.Profile},
		{"GET", "/user/index", controllers.UserController.Index},
		{"GET", "/user/profile", controllers.UserController.Profile},
	}
}

func main() {
	//test1()

	app := WebApp{}
	app.init()
	app.start()
}

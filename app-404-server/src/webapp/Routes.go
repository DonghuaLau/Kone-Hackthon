package main

import (
	//"fmt"
	"net/http"
)

type Router struct {
	_method  string
	_uri     string
	_handler func(http.ResponseWriter, *http.Request)
}

/*
	app._routers = []Router{
		{"GET", "/index", controllers.HomeController.Index},
		{"GET", "/profile", controllers.HomeController.Profile},
		{"GET", "/user/index", controllers.UserController.Index},
		{"GET", "/user/profile", controllers.UserController.Profile},
	}
*/

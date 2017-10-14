package controllers

import (
	"fmt"
	"net/http"
)

var HomeController _HomeController

type _HomeController struct{}

func (ctl *_HomeController) Index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "HomeController index\n")
}

func (ctl *_HomeController) Profile(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "HomeController profile\n")
}

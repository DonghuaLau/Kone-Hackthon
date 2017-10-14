package controllers

import (
	"fmt"
	"net/http"
)

var UserController _UserController

type _UserController struct{}

func (ctl *_UserController) Index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "UserController index\n")
}

func (ctl *_UserController) Profile(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "UserController profile\n")
}

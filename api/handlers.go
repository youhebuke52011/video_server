package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w,"hello handlers")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("user_name")
	io.WriteString(w, name)
}
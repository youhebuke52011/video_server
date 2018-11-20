package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateHandler)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	handlers := RegisterHandlers()
	http.ListenAndServe(":8000", handlers)
}

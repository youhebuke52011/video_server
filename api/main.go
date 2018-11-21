package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	handler := middleWareHandler{}
	handler.r = r
	return handler
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//

	m.r.ServeHTTP(w, req)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateHandler)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	handler := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", handler)
}

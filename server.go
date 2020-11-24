package main

import (
	"github.com/dimfeld/httptreemux"
	"net/http"
	HANDLERS "./handlers"
)

func main() {
	router := httptreemux.New()
	bookGroup := router.NewGroup("/books")
	bookGroup.GET("/", HANDLERS.GetBooks)
	bookGroup.POST("/post", HANDLERS.PostBook)
	http.ListenAndServe(":8080", router)
}
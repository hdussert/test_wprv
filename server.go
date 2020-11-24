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
	bookGroup.PUT("/update", HANDLERS.UpdateBook)
	bookGroup.DELETE("/delete", HANDLERS.DeleteBook)
	bookGroup.DELETE("/deletex", HANDLERS.DeleteBooks)
	http.ListenAndServe(":8080", router)
}
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
	bookGroup.GET("/filter", HANDLERS.FilterBooks) // "/filter?search=string&date_start=2000-01-30&date_end=2010-01-30"
	
	bookGroup.POST("/", HANDLERS.PostBook)
	bookGroup.PUT("/", HANDLERS.UpdateBook)
	bookGroup.DELETE("/", HANDLERS.DeleteBooks)
	
	http.ListenAndServe(":8080", router)
}
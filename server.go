package main

import (
	"github.com/dimfeld/httptreemux"
	HANDLERS "./handlers"
)

func main() {
	router := httptreemux.New()
	bookGroup := router.NewGroup("/books")
	bookGroup.GET("/", HANDLERS.GetBooks)
}
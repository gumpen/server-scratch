package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/gumpen/server-scratch/handler"
)

func main() {
	router := httprouter.New()
	router.GET("/users", handler.UserIndex)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

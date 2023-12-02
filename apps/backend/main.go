package main

import (
	"log"
	"net/http"
)

const port string = ":8880"

func main() {
	log.Println("starting backend...")

	router := routes()
	log.Println("running on port: ", port)
	http.ListenAndServe(port, router)
}

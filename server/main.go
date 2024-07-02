package main

import (
	"log"
	"net/http"
	"server/pkg/routes"
)

func main() {
	if err := http.ListenAndServe("localhost:8080", routes.Serve); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

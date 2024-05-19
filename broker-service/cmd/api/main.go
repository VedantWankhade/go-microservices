package main

import (
	"log"
	"net/http"
)

const port = ":80"

type Config struct{}

func main() {
	app := Config{}
	log.Println("Starting broker service on", port)
	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}
	log.Fatal(srv.ListenAndServe())
}

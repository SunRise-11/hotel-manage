package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

const (
	DEFAULT_PORT = 3000
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	store := NewSessionStore()
	initDb()
	initReaper(store)

	addr := fmt.Sprintf(":%v", DEFAULT_PORT)
	mainRouter := handlers.LoggingHandler(os.Stdout, createRouter(store))

	// TODO: Run in goroutine with signal handling to not block
	// https://github.com/gorilla/mux
	log.Println("Running server on", addr)
	log.Fatal(http.ListenAndServe(addr, mainRouter))
}

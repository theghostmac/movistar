// Package main initializes all components of the rating service and starts an HTTP handler.
package main

import (
	"log"
	"movistar/rating/internal/controller/rating"
	httpHandler "movistar/rating/internal/handler/http"
	"movistar/rating/internal/repository/memory"
	"net/http"
)

func main() {
	log.Println("Starting the rating service...")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}

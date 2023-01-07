package main

import (
	"log"
	"net/http"

	"movistar/metadata/internal/controller/metadata"
	httpHandler "movistar/metadata/internal/handler/http"
	"movistar/metadata/internal/repository/memory"
)

// main function will initialize all structures of the service and start the http PI handler.

func main() {
	log.Println("Starting the movie metadata service...")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

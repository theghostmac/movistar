package main

import (
	"log"
	"net/http"

	"movistar/movie/internal/controller/movie"
	metadataGateway "movistar/movie/internal/gateway/metadata/http"
	ratingGateway "movistar/movie/internal/gateway/rating/http"
	httpHandler "movistar/movie/internal/handler/http"
)

func main() {
	log.Println("Starting the movistar...")
	metadataGatewayVar := metadataGateway.New("localhost:8081")
	ratingGatewayVar := ratingGateway.New("localhost:8082")
	ctrl := movie.New(ratingGatewayVar, metadataGatewayVar)
	h := httpHandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/0xAckerMan/movieapp-ms/movie/internal/controller/movie"
	metadatagateway "github.com/0xAckerMan/movieapp-ms/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/0xAckerMan/movieapp-ms/movie/internal/gateway/rating/http"
	httphandler "github.com/0xAckerMan/movieapp-ms/movie/internal/handler/http"
)

func main(){
    log.Printf("Starting the movie service")
    metadatagateway := metadatagateway.New("localhost:8081")
    ratinggateway := ratinggateway.New("localhost:8082")
    ctrl := movie.New(ratinggateway,metadatagateway)
    h := httphandler.New(ctrl)
    http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
    if err := http.ListenAndServe(":8083",nil); err != nil{
        panic(err)
    }
}

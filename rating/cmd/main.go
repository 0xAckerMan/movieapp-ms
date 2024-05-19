package main

import (
	"log"
	"net/http"

	"github.com/0xAckerMan/movieapp-ms/rating/internal/controller/rating"
	httphandler "github.com/0xAckerMan/movieapp-ms/rating/internal/handler/http"
	"github.com/0xAckerMan/movieapp-ms/rating/internal/repository/memory"
)

func main(){
    log.Printf("Starting the rating service")
    repo := memory.New()
    ctrl := rating.New(repo)
    h := httphandler.New(ctrl)
    http.Handle("/rating", http.HandlerFunc(h.Handler))
    if err := http.ListenAndServe(":8082", nil); err != nil{
        panic(err)
    }
}

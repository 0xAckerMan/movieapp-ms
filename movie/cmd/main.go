package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/0xAckerMan/movieapp-ms/movie/internal/controller/movie"
	metadatagateway "github.com/0xAckerMan/movieapp-ms/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/0xAckerMan/movieapp-ms/movie/internal/gateway/rating/http"
	httphandler "github.com/0xAckerMan/movieapp-ms/movie/internal/handler/http"
	"github.com/0xAckerMan/movieapp-ms/pkg/discovery"
	"github.com/0xAckerMan/movieapp-ms/pkg/discovery/consul"
)

const serviceName = "movie"

func main() {
    var port int
    flag.IntVar(&port, "port", 8083, "API handler port")
    flag.Parse()
    
	log.Printf("Starting the movie service")

    registry, err := consul.NewRegistry("localhost:8500")
    if err != nil{
        panic(err)
    }
    ctx := context.Background()
    instanceID := discovery.GenerateInstanceID(serviceName)
    if err := registry.Register(ctx,serviceName,instanceID, fmt.Sprintf("localhost:%d", port)); err != nil{
        panic(err)
    }

    go func(){
        for{
            if err := registry.ReportHealthState(instanceID, serviceName); err != nil{
                log.Println("Failed to report health state: "+ err.Error())
            }
            time.Sleep(1*time.Second)
        }
    }()
    defer registry.Deregister(ctx,instanceID,serviceName)
	metadatagateway := metadatagateway.New(registry)
	ratinggateway := ratinggateway.New(registry)
	svc := movie.New(ratinggateway, metadatagateway)
	h := httphandler.New(svc)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}

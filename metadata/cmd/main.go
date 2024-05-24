package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/0xAckerMan/movieapp-ms/metadata/internal/controller/metadata"
	httphandler "github.com/0xAckerMan/movieapp-ms/metadata/internal/handler/http"
	"github.com/0xAckerMan/movieapp-ms/metadata/internal/repository/memory"
	"github.com/0xAckerMan/movieapp-ms/pkg/discovery"
	"github.com/0xAckerMan/movieapp-ms/pkg/discovery/consul"
)

const serviceName = "metadata"

func main() {
    var port int
    flag.IntVar(&port, "port", 8081, "API handler port")
    flag.Parse()

	log.Println("Starting the movie metadata service", port)

    registry, err := consul.NewRegistry("localhost:8500")
    if err != nil{
        panic(err)
    }
    ctx := context.Background()

    instanceID := discovery.GenerateInstanceID(serviceName)

    if err := registry.Register(ctx,serviceName,instanceID,fmt.Sprintf("localhost:%d", port)); err != nil{
        panic(err)
    }
    go func() {
        for {
            if err := registry.ReportHealthState(instanceID, serviceName); err != nil{
                log.Println("failed to report health state: " + err.Error())
            }
            time.Sleep(1*time.Second)
        }
    }()
    defer registry.Deregister(ctx,instanceID, serviceName)

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

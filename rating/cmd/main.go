package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/0xAckerMan/movieapp-ms/pkg/discovery"
	"github.com/0xAckerMan/movieapp-ms/pkg/discovery/consul"
	"github.com/0xAckerMan/movieapp-ms/rating/internal/controller/rating"
	grpchandler "github.com/0xAckerMan/movieapp-ms/rating/internal/handler/grpc"
	"github.com/0xAckerMan/movieapp-ms/rating/internal/repository/memory"
	"github.com/0xAckerMan/movieapp-ms/src/gen"
	"google.golang.org/grpc"
)
const serviceName = "rating"

func main() {
    var port int
    flag.IntVar(&port, "port", 8082, "API handler port")
    flag.Parse()
	log.Printf("Starting the rating service")
    registry, err := consul.NewRegistry("localhost:8500")
    if err != nil{
        panic(err)
    }
    ctx := context.Background()
    instanceID := discovery.GenerateInstanceID(serviceName)
    if err := registry.Register(ctx,serviceName,instanceID, fmt.Sprintf("localhost:%d", port)); err != nil {
        panic(err)
    }

    go func(){
        for {
            if err := registry.ReportHealthState(instanceID,serviceName); err != nil{
                log.Println("failed to report health state" + err.Error())
            }
            time.Sleep(1 * time.Second)
        }
    }()
    defer registry.Deregister(ctx,instanceID,serviceName)

	repo := memory.New()
	ctrl := rating.New(repo)
	h := grpchandler.New(ctrl)
    lis, err := net.Listen("tcp", "localhost:8082")
    if err != nil{
        log.Fatalf("Failed to listen: %v", err)
    }
    srv := grpc.NewServer()
    gen.RegisterRatingServiceServer(srv,h)
    srv.Serve(lis)
}

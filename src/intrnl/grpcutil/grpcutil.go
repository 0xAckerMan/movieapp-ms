package grpcutil

import (
	"context"
	"math/rand"

	"github.com/0xAckerMan/movieapp-ms/pkg/discovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(ctx context.Context, serviceName string, registy discovery.Registry)(*grpc.ClientConn, error){
    addrs, err := registy.ServiceAddresses(ctx, serviceName)
    if err != nil{
        return nil, err
    }
    return grpc.NewClient(addrs[rand.Intn(len(addrs))], grpc.WithTransportCredentials(insecure.NewCredentials()))
}

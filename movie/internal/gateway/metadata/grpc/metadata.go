package grpc

import (
	"context"

	model "github.com/0xAckerMan/movieapp-ms/metadata/pkg"
    models "github.com/0xAckerMan/movieapp-ms/metadata/pkg/model"
	"github.com/0xAckerMan/movieapp-ms/pkg/discovery"
	"github.com/0xAckerMan/movieapp-ms/src/gen"
	grpcutil "github.com/0xAckerMan/movieapp-ms/src/intrnl/grpcutil"
)

type Gateway struct{
    registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway{
    return &Gateway{registry}
}

func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error){
    conn, err := grpcutil.ServiceConnection(ctx, "metadata", g.registry)
    if err != nil{
        return nil, err
    }
    defer conn.Close()

    client := gen.NewMetadataServiceClient(conn)
    resp, err := client.GetMetadata(ctx,&gen.GetMetadataRequest{MovieId: id})
    if err != nil{
        return nil, err
    }

    return models.MetadataFromProto(resp.Metadata), nil
}

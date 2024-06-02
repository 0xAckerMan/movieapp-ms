package grpc

import (
	"context"

	"github.com/0xAckerMan/movieapp-ms/pkg/discovery"
	model "github.com/0xAckerMan/movieapp-ms/rating/pkg"
	"github.com/0xAckerMan/movieapp-ms/src/gen"
	"github.com/0xAckerMan/movieapp-ms/src/intrnl/grpcutil"
)

type Gateway struct{
    registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway{
    return &Gateway{registry}
}

// GetAggregatedRating returns the aggregated rating for a
// record or ErrNotFound if there are no ratings for it.

func (g *Gateway) GetAggregatedRating (ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error){
    conn, err := grpcutil.ServiceConnection(ctx,"rating", g.registry)
    if err != nil{
        return 0, err
    }

    defer conn.Close()

    client := gen.NewRatingServiceClient(conn)
    resp, err := client.GetAggregatedRating(ctx,&gen.GetAggregatedRatingRequest{RecordId: string(recordID), RecordType: string(recordType)})
    if err != nil{
        return 0, err
    }

    return resp.RatingValue, nil
}

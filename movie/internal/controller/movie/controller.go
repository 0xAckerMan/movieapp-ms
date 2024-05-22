package movie

import (
	"context"
	"errors"

	metadatamodel "github.com/0xAckerMan/movieapp-ms/metadata/pkg"
	"github.com/0xAckerMan/movieapp-ms/movie/internal/gateway"
	"github.com/0xAckerMan/movieapp-ms/movie/pkg/model"
	ratingmodel "github.com/0xAckerMan/movieapp-ms/rating/pkg"
)

var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface{
    GetAggregateRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
    PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface{
    Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

type Controller struct{
    ratingGateway ratingGateway
    metadataGateway metadataGateway
}

func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller{
    return &Controller{ratingGateway, metadataGateway}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error){
    metadata, err := c.metadataGateway.Get(ctx,id)
    if err != nil && errors.Is(err, gateway.ErrNotFound){
        return nil, ErrNotFound
    } else if err != nil{
        return nil, err
    }

    details:= &model.MovieDetails{
        Metadata: *metadata,
    }

    rating, err:= c.ratingGateway.GetAggregateRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)
    if err != nil && errors.Is(err, gateway.ErrNotFound){
        // it is safe to proceed not having a rating at the momment
    } else if err != nil{
        return nil, err
    } else {
        details.Rating = &rating
    }
    return details, err
}

package rating

import (
	"context"
	"errors"

	"github.com/0xAckerMan/movieapp-ms/rating/internal/repository"
	model "github.com/0xAckerMan/movieapp-ms/rating/pkg"
)

// Defines the error for a not found record
var ErrNotFound = errors.New("Rating not found for a record")

type ratingRepository interface {
	Get(ctx context.Context, recordId model.RecordID, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error
}

// Controller defines a rating service controller
type Controller struct {
	repo ratingRepository
    ingester ratingIngester
}

type ratingIngester interface{
    Ingest(ctx context.Context) (chan model.Rating, error)
}

func New(repo ratingRepository, ingester ratingIngester) *Controller {
	return &Controller{repo, ingester}
}

func (c *Controller) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordID, recordType)
	if err != nil && err == repository.ErrNotFound {
		return 0, ErrNotFound
	} else if err != nil {
		return 0, err
	}

	sum := float64(0)

	for _, r := range ratings {
		sum += float64(r.Value)
	}
	return sum / float64(len(ratings)), nil
}

func (c *Controller) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	return c.repo.Put(ctx, recordID, recordType, rating)
}

//StartIngestion starts the ingestion of rating events
func (s *Controller) StartIngestion(ctx context.Context) error {
    ch, err := s.ingester.Ingest(ctx)
    if err != nil{
        return err
    }

    for e := range ch{
        if err := s.PutRating(ctx, model.RecordID(e.RecordID), model.RecordType(e.RecordType),&model.Rating{UserID: e.UserID,Value: e.Value}); err != nil{
            return err
        }
    }
    return err
}

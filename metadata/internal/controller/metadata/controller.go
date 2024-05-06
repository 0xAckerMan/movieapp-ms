package metadata

import (
	"context"
	"errors"

	"github.com/0xAckerMan/movieapp-ms/metadata/internal/repository"
	model "github.com/0xAckerMan/movieapp-ms/metadata/pkg"
)

// ErrNotFound is returned when the requested record is not found
var ErrNotFound = errors.New("not found")

type MetadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

//Controller defines a metadata service controller
type Controller struct{
    repo MetadataRepository
}

//New creates a metadata service controller
func New(repo MetadataRepository) *Controller{
    return &Controller{repo}
}

//Get returns movie metadata by id
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error){
    res, err := c.repo.Get(ctx,id)
    if err != nil && errors.Is(err, repository.ErrNotFound){
        return nil, ErrNotFound
    }
    return res, nil
}

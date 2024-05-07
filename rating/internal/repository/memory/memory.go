package memory

import (
	"context"

	"github.com/0xAckerMan/movieapp-ms/rating/internal/repository"
	model "github.com/0xAckerMan/movieapp-ms/rating/pkg"
)

//Repository defines a rating repository
type Repository struct{
    data map[model.RecordType] map[model.RecordID] []model.Rating
}

// Creates a new memory repository
func New() *Repository{
    return &Repository{map[model.RecordType]map[model.RecordID][]model.Rating{}}
}

// retrieves all ratings for a given record
func (r *Repository) Get (ctx context.Context, recordId model.RecordID, recordType model.RecordType) ([]model.Rating, error){
    if _, ok := r.data[recordType]; !ok {
        return nil, repository.ErrNotFound 
    }

    if ratings, ok := r.data[recordType][recordId]; !ok || len(ratings) == 0 {
        return nil, repository.ErrNotFound
    }

    return r.data[recordType][recordId], nil
}

// Put adds a rating for a given record
func (r *Repository) Put(ctx context.Context, recordId model.RecordID, recordType model.RecordType, rating *model.Rating) error {
    if _, ok := r.data[recordType]; !ok {
        r.data[recordType] = map[model.RecordID][]model.Rating{}
    }

    r.data[recordType][recordId] = append(r.data[recordType][recordId], *rating)

    return nil
}

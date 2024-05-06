package memory

import (
	"context"
	"sync"

	"github.com/0xAckerMan/movieapp-ms/metadata/internal/repository"
	model "github.com/0xAckerMan/movieapp-ms/metadata/pkg"
)

type Repository struct{
    sync.RWMutex
    data map[string]*model.Metadata
}

func New() *Repository{
    return &Repository{data: map[string]*model.Metadata{}}
}

func (r *Repository) Get (_ context.Context, id string) (*model.Metadata, error){
    r.RLock()
    defer r.RUnlock()

    m, ok := r.data[id]
    if !ok{
        return nil, repository.ErrNotFound
    }

    return m, nil
}

func (r *Repository) Put (_ context.Context, id string, metadata *model.Metadata) error {
    r.Lock()
    defer r.Unlock()

    r.data[id] = metadata
    return nil
}

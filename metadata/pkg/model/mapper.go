package model

import (
	model "github.com/0xAckerMan/movieapp-ms/metadata/pkg"
	"github.com/0xAckerMan/movieapp-ms/src/gen"
)

func MetadataToProto(m *model.Metadata) *gen.Metadata{
    return &gen.Metadata{
        Id: m.ID,
        Title: m.Title,
        Description: m.Description,
        Director: m.Director,
    }
} 

func MetadataFromProto (m *gen.Metadata) *model.Metadata{
    return &model.Metadata{
        ID: m.Id,
        Title: m.Title,
        Description: m.Description,
        Director: m.Director,
    }
}


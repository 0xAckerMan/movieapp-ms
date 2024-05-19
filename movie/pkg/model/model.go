package model

import (
	model "github.com/0xAckerMan/movieapp-ms/metadata/pkg"
)

type MovieDetails struct{
    Rating *float64 `json:"rating,omitEmpty"`
    Metadata model.Metadata `json:"metadata"`
}

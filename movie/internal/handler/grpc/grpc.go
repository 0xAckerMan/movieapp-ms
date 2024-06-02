package grpc

import (
	"context"
	"errors"

	"github.com/0xAckerMan/movieapp-ms/metadata/pkg/model"
	"github.com/0xAckerMan/movieapp-ms/movie/internal/controller/movie"
	"github.com/0xAckerMan/movieapp-ms/src/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct{
    gen.UnimplementedMovieServiceServer
    ctrl *movie.Controller
}

func New(ctrl *movie.Controller) *Handler{
    return &Handler{ctrl: ctrl}
}

func (h *Handler) GetMovieDetails(ctx context.Context, req *gen.GetMovieDetailsRequest) (*gen.GetMovieDetailsResponse, error){
    if req == nil || req.MovieId == "" {
        return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
    }

    m, err := h.ctrl.Get(ctx,req.MovieId)
    if err != nil && errors.Is(err, movie.ErrNotFound){
        return nil, status.Errorf(codes.NotFound, err.Error())
    } else if err != nil{
        return nil, status.Errorf(codes.Internal, err.Error())
    }

    return &gen.GetMovieDetailsResponse{
        MovieDetails: &gen.MovieDetails{
            Metadata: model.MetadataToProto(&m.Metadata),
            Rating: float32(*m.Rating),
        },
    }, nil
}

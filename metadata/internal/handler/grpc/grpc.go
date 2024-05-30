package grpc

import (
	"context"
	"errors"

	"github.com/0xAckerMan/movieapp-ms/metadata/internal/controller/metadata"
	"github.com/0xAckerMan/movieapp-ms/metadata/internal/repository"
	"github.com/0xAckerMan/movieapp-ms/metadata/pkg/model"
	"github.com/0xAckerMan/movieapp-ms/src/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedMetadataServiceServer
	ctrl *metadata.Controller
}

func New(ctrl *metadata.Controller) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}

func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	if req == nil || req.MovieId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.MovieId)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

    return &gen.GetMetadataResponse{
        Metadata: model.MetadataToProto(m),
    }, nil
}

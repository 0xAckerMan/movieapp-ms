package grpc

import (
	"context"
	"errors"

	"github.com/0xAckerMan/movieapp-ms/rating/internal/controller/rating"
	model "github.com/0xAckerMan/movieapp-ms/rating/pkg"
	"github.com/0xAckerMan/movieapp-ms/src/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct{
    gen.UnimplementedRatingServiceServer
    ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler{
    return &Handler{ctrl: ctrl}
}

func (h *Handler) GetAggregatedRating(ctx context.Context, req *gen.GetAggregatedRatingRequest) (*gen.GetAggregatedRatingResponse, error){
    if req == nil || req.RecordId == "" || req.RecordType == "" {
        return nil, status.Errorf(codes.InvalidArgument, "nil req or record id")
    }
    v, err := h.ctrl.GetAggregatedRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType))
    if err != nil && errors.Is(err,rating.ErrNotFound){
        return nil, status.Errorf(codes.NotFound, err.Error())
    } else if err != nil{
        return nil, status.Errorf(codes.Internal, err.Error())
    }

    return &gen.GetAggregatedRatingResponse{RatingValue: v}, nil
}

func (h *Handler) PutRating (ctx context.Context,req *gen.PutRatingRequest) (*gen.PutRatingResponse, error){
    if req == nil || req.UserId == "" || req.RecordId == "" {
        return nil, status.Errorf(codes.InvalidArgument, "nil req or empty UserId or RecordId")
    }
    if err := h.ctrl.PutRating(ctx,model.RecordID(req.RecordId), model.RecordType(req.RecordType), &model.Rating{UserID: model.UserID(req.UserId), Value: model.RatingValue(req.RatingValue)}); err != nil{
        return nil, err
    }
    return &gen.PutRatingResponse{}, nil
}

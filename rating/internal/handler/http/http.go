package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/0xAckerMan/movieapp-ms/rating/internal/controller/rating"
	model "github.com/0xAckerMan/movieapp-ms/rating/pkg"
)

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	recordID := model.RecordID(r.FormValue("id"))
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordType := model.RecordType(r.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(r.Context(), recordID, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Encode response error: %v \n", err)
		}
    case http.MethodPut:
        userID := model.UserID(r.FormValue("userId"))
        v, err := strconv.ParseFloat(r.FormValue("value"),64)
        if err != nil{
            w.WriteHeader(http.StatusBadRequest)
            return
        }
    if err := h.ctrl.PutRating(context.Background(),recordID,recordType, &model.Rating{UserID: userID, Value: model.RatingValue(v)}); err != nil{
            log.Printf("Repository put error: %v \n", err)
            w.WriteHeader(http.StatusInternalServerError)
        }
    default:
    w.WriteHeader(http.StatusBadRequest)
	}

}

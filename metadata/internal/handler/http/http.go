package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/0xAckerMan/movieapp-ms/metadata/internal/controller/metadata"
	"github.com/0xAckerMan/movieapp-ms/metadata/internal/repository"
)

type Handler struct{
    ctrl *metadata.Controller
}

func New(ctrl *metadata.Controller) *Handler{
    return &Handler{ctrl}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request){
    id := req.FormValue("id")
    if id == ""{
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    ctx := req.Context()
    m, err := h.ctrl.Get(ctx,id)
    if err != nil && errors.Is(err, repository.ErrNotFound){
        w.WriteHeader(http.StatusNotFound)
        return
    } else if err != nil{
        log.Printf("repository get err %v \n", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(m); err != nil{
        log.Printf("Response encode error %v \n", err)
    }
}

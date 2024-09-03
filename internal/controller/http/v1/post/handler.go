package post

import (
	"github.com/skripov-ds-ai/highload_course/internal/generated"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// (POST /post/create)
func (h *Handler) PostPostCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PUT /post/delete/{id})
func (h *Handler) PutPostDeleteID(w http.ResponseWriter, r *http.Request, id generated.PostID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /post/feed)
func (h *Handler) GetPostFeed(w http.ResponseWriter, r *http.Request, params generated.GetPostFeedParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /post/get/{id})
func (h *Handler) GetPostGetID(w http.ResponseWriter, r *http.Request, id generated.PostID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PUT /post/update)
func (h *Handler) PutPostUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

package dialog

import (
	"github.com/skripov-ds-ai/highload_course/internal/generated"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// (GET /dialog/{user_id}/list)
func (h *Handler) GetDialogUserIDList(w http.ResponseWriter, r *http.Request, userId generated.UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /dialog/{user_id}/send)
func (h *Handler) PostDialogUserIDSend(w http.ResponseWriter, r *http.Request, userId generated.UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

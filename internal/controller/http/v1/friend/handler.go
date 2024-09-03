package friend

import (
	"github.com/skripov-ds-ai/highload_course/internal/generated"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// (PUT /friend/delete/{user_id})
func (h *Handler) PutFriendDeleteUserID(w http.ResponseWriter, r *http.Request, userId generated.UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PUT /friend/set/{user_id})
func (h *Handler) PutFriendSetUserID(w http.ResponseWriter, r *http.Request, userId generated.UserID) {
	w.WriteHeader(http.StatusNotImplemented)
}

package login

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	models "github.com/skripov-ds-ai/highload_course/internal/entity"
	"github.com/skripov-ds-ai/highload_course/internal/generated"
	"net/http"
)

type AuthService interface {
	Login(ctx context.Context, userID, password string) (string, error)
}

type Handler struct {
	authService AuthService
	validate    *validator.Validate
}

func NewHandler(authSrv AuthService) *Handler {
	return &Handler{
		authService: authSrv,
		validate:    validator.New(),
	}
}

// (POST /login)
func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := middleware.GetReqID(ctx)

	var body generated.PostLoginJSONBody
	err := render.DecodeJSON(r.Body, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.validate.Struct(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(ctx, body.ID, body.Password)
	if errors.Is(err, models.ErrNotFound) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if errors.Is(err, models.ErrWrongPassword) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, generated.N5Xx{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			RequestID: &requestID,
		})
		return
	}

	render.JSON(w, r, models.UserToken{
		Token: token,
	})
}

package user

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	models "github.com/skripov-ds-ai/highload_course/internal/entity"
	"github.com/skripov-ds-ai/highload_course/internal/generated"
	"go.uber.org/zap"
	"net/http"

	"github.com/go-chi/render"
)

type UserService interface {
	Get(ctx context.Context, userID string) (models.User, error)
	Register(ctx context.Context, params models.CreateUserParams) (string, error)
	ListByPrefixFirstNameSecondName(ctx context.Context, firstName, secondName string) (models.Users, error)
}

type Handler struct {
	userService UserService
	logger      *zap.Logger
}

func NewHandler(userSrv UserService, logger *zap.Logger) *Handler {
	return &Handler{
		userService: userSrv,
		logger:      logger,
	}
}

// (GET /user/get/{id})
func (h *Handler) GetUserGetId(w http.ResponseWriter, r *http.Request, id generated.UserID) {
	//w.WriteHeader(http.StatusNotImplemented)
	_, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	requestID := middleware.GetReqID(ctx)

	u, err := h.userService.Get(ctx, id)
	h.logger.Error("while get", zap.Error(err))
	if errors.Is(err, models.ErrNotFound) {
		// TODO: ErrRenderer
		//func ErrRender(err error) render.Renderer {
		//	return &ErrResponse{
		//	Err:            err,
		//	HTTPStatusCode: 422,
		//	StatusText:     "Error rendering response.",
		//	ErrorText:      err.Error(),
		//}
		//}
		//
		//var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

		//render.Render(w, r, ErrRender(err))
		//render.Status(r, http.StatusInternalServerError)

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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

	render.JSON(w, r, u.ToModel())
}

// (POST /user/register)
func (h *Handler) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := middleware.GetReqID(ctx)

	var body generated.PostUserRegisterJSONBody
	err := render.DecodeJSON(r.Body, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := h.userService.Register(ctx, models.NewCreateUserParams(body))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, generated.N5Xx{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			RequestID: &requestID,
		})
		return
	}

	render.JSON(w, r, models.RegisterUserResponse{UserID: userID})
}

// (GET /user/search)
func (h *Handler) GetUserSearch(w http.ResponseWriter, r *http.Request, params generated.GetUserSearchParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

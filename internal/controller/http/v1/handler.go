package v1

import (
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/dialog"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/friend"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/login"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/post"
	"github.com/skripov-ds-ai/highload_course/internal/controller/http/v1/user"
	"github.com/skripov-ds-ai/highload_course/internal/generated"
	"net/http"
)

type Handler struct {
	dialog *dialog.Handler
	friend *friend.Handler
	login  *login.Handler
	post   *post.Handler
	user   *user.Handler
}

func NewHandler(
	dialog *dialog.Handler,
	friend *friend.Handler,
	login *login.Handler,
	post *post.Handler,
	user *user.Handler) *Handler {
	return &Handler{
		dialog: dialog,
		friend: friend,
		login:  login,
		post:   post,
		user:   user,
	}
}

// (GET /dialog/{user_id}/list)
func (h *Handler) GetDialogUserIDList(w http.ResponseWriter, r *http.Request, userId generated.UserID) {
	h.dialog.GetDialogUserIDList(w, r, userId)
}

// (POST /dialog/{user_id}/send)
func (h *Handler) PostDialogUserIDSend(w http.ResponseWriter, r *http.Request, userId generated.UserID) {
	h.dialog.PostDialogUserIDSend(w, r, userId)
}

// (PUT /friend/delete/{user_id})
func (h *Handler) PutFriendDeleteUserID(w http.ResponseWriter, r *http.Request, userId generated.UserID) {
	h.friend.PutFriendDeleteUserID(w, r, userId)
}

// (PUT /friend/set/{user_id})
func (h *Handler) PutFriendSetUserID(w http.ResponseWriter, r *http.Request, userId generated.UserID) {
	h.friend.PutFriendSetUserID(w, r, userId)
}

// (POST /login)
func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
	h.login.PostLogin(w, r)
}

// (POST /post/create)
func (h *Handler) PostPostCreate(w http.ResponseWriter, r *http.Request) {
	h.post.PostPostCreate(w, r)
}

// (PUT /post/delete/{id})
func (h *Handler) PutPostDeleteID(w http.ResponseWriter, r *http.Request, id generated.PostID) {
	h.post.PutPostDeleteID(w, r, id)
}

// (GET /post/feed)
func (h *Handler) GetPostFeed(w http.ResponseWriter, r *http.Request, params generated.GetPostFeedParams) {
	h.post.GetPostFeed(w, r, params)
}

// (GET /post/get/{id})
func (h *Handler) GetPostGetID(w http.ResponseWriter, r *http.Request, id generated.PostID) {
	h.post.GetPostGetID(w, r, id)
}

// (PUT /post/update)
func (h *Handler) PutPostUpdate(w http.ResponseWriter, r *http.Request) {
	h.post.PutPostUpdate(w, r)
}

// (GET /user/get/{id})
func (h *Handler) GetUserGetID(w http.ResponseWriter, r *http.Request, id generated.UserID) {
	h.user.GetUserGetId(w, r, id)
}

// (POST /user/register)
func (h *Handler) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	h.user.PostUserRegister(w, r)
}

// (GET /user/search)
func (h *Handler) GetUserSearch(w http.ResponseWriter, r *http.Request, params generated.GetUserSearchParams) {
	h.user.GetUserSearch(w, r, params)
}

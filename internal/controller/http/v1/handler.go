package v1

import (
	"highload_course/internal/generated"
)

type Handler struct {
	apiHandler generated.ServerInterface
}

func (h Handler) SetupHandlers() {
	//h.apiHandler.GetPostFeed = func(w http.ResponseWriter, r *http.Request, params generated.GetPostFeedParams) {
	//
	//}
}

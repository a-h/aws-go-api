package user

import (
	"net/http"

	"github.com/a-h/aws-go-api/db"
	"github.com/a-h/aws-go-api/respond"
)

func NewHandler(us db.UserStore, vars func(r *http.Request) map[string]string) Handler {
	return Handler{
		UserStore: us,
		Vars:      vars,
	}
}

type Handler struct {
	UserStore db.UserStore
	Vars      func(r *http.Request) map[string]string
}

func (h Handler) DetailsGet(w http.ResponseWriter, r *http.Request) {
	vars := h.Vars(r)
	id, ok := vars["id"]
	if !ok {
		respond.WithError(w, http.StatusNotFound, "id parameter missing")
		return
	}
	details, err := h.UserStore.GetDetails(id)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "failed to get user")
		return
	}
	respond.WithJSON(w, http.StatusOK, details)
}

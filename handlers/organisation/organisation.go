package organisation

import (
	"net/http"

	"github.com/a-h/aws-go-api/db"
	"github.com/a-h/aws-go-api/respond"
)

func NewHandler(os db.OrganisationStore, vars func(r *http.Request) map[string]string) Handler {
	return Handler{
		OrganisationStore: os,
		Vars:              vars,
	}
}

type Handler struct {
	OrganisationStore db.OrganisationStore
	Vars              func(r *http.Request) map[string]string
}

func (h Handler) DetailsGet(w http.ResponseWriter, r *http.Request) {
	vars := h.Vars(r)
	id, ok := vars["id"]
	if !ok {
		respond.WithError(w, http.StatusNotFound, "id parameter missing")
		return
	}
	details, err := h.OrganisationStore.GetDetails(id)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "failed to get organisation")
		return
	}
	respond.WithJSON(w, http.StatusOK, details)
}

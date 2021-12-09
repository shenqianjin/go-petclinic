package handler

import (
	"net/http"
	"petclinic/internal/support"
)

func NewPetForm(ctx *support.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doHandlerInternal(w, r, "../../web/templates/pets/create_pet.html", nil)
	}
}

func NewVisitForm(ctx *support.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doHandlerInternal(w, r, "../../web/templates/pets/create_visit.html", nil)
	}
}

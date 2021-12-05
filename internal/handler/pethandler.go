package handler

import (
	"net/http"
	"petclinic/internal/service"
)

func NewPetForm(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doHandlerInternal(w, r, "../web/templates/pets/create_pet.html", nil)
	}
}

func NewVisitForm(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doHandlerInternal(w, r, "../web/templates/pets/create_visit.html", nil)
	}
}


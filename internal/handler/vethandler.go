package handler

import (
	"fmt"
	"net/http"
	"petclinic/internal/biz"
	"petclinic/internal/support"
)

func NewVetForm(ctx *support.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doHandlerInternal(w, r, "../../web/templates/vets/create_vet.html", nil)
	}
}

func GetVets(ctx *support.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vetList, err := biz.GetPetBiz().Vets()
		if err != nil {
			http.Error(w, fmt.Sprintf("unable to get vets: %q\n", err),
				http.StatusInternalServerError)
			return
		}
		doHandlerInternal(w, r, "../../web/templates/vets/vet_list.html", vetList)
	}
}

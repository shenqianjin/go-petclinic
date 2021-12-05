package handler

import (
	"fmt"
	"net/http"
	"petclinic/internal/biz"
	"petclinic/internal/model"
	"petclinic/internal/service"
	"strconv"
)

func NewOwnerForm(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Post接口
			// validated
			if r.FormValue("firstName") == "" {
				http.Error(w, fmt.Sprintf("firstName can not be empty"), http.StatusInternalServerError)
				return
			} else if r.FormValue("lastName") == "" {
				http.Error(w, fmt.Sprintf("lastName can not be empty"), http.StatusInternalServerError)
			} else if r.FormValue("address") == "" {
				http.Error(w, fmt.Sprintf("address can not be empty"), http.StatusInternalServerError)
			} else if r.FormValue("city") == "" {
				http.Error(w, fmt.Sprintf("city can not be empty"), http.StatusInternalServerError)
			} else if r.FormValue("telephone") == "" {
				http.Error(w, fmt.Sprintf("telephone can not be empty"), http.StatusInternalServerError)
			} else {
				owner := model.Owner{
					FirstName: r.FormValue("firstName"),
					LastName:  r.FormValue("lastName"),
					Address:   r.FormValue("address"),
					City:      r.FormValue("city"),
					Telephone: r.FormValue("telephone"),
				}
				id := biz.GetOwnerBiz().AddOwner(owner)
				http.Redirect(w, r, "/owners/detail?id="+strconv.FormatInt(id, 10), 302)
			}
		} else if r.Method == http.MethodGet {
			// page
			doHandlerInternal(w, r, "../web/templates/owners/create_owner.html", nil)
		}
	}
}

func FindOwner(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doHandlerInternal(w, r, "../web/templates/owners/find_owner.html", nil)
	}
}

func GetOwners(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get request param
		lastName := r.URL.Query().Get("lastName")
		// call service
		ownersList, err := biz.GetOwnerBiz().Owners(lastName)
		// validate after call service
		if err != nil {
			http.Error(w, fmt.Sprintf("unable to get vets: %q\n", err),
				http.StatusInternalServerError)
			return
		}
		// render view
		doHandlerInternal(w, r, "../web/templates/owners/owners_list.html", ownersList)
	}
}

func GetOwnerDetail(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get request param
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.ParseInt(idStr, 10, 64)
		// call service
		owner, err := biz.GetOwnerBiz().FindOwner(id)
		// validate after call service
		if err != nil {
			http.Error(w, fmt.Sprintf("unable to get vets: %q\n", owner),
				http.StatusInternalServerError)
			return
		}
		// render view
		doHandlerInternal(w, r, "../web/templates/owners/owner_detail.html", owner)
	}
}


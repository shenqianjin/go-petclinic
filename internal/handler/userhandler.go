package handler

import (
	"net/http"
	"petclinic/internal/model/dto"

	"github.com/tal-tech/go-zero/rest/httpx"
	"petclinic/internal/biz"
	"petclinic/internal/service"
)

func CreateUserHandler(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := biz.NewCreateUserLogic(r.Context(), ctx)
		err := l.CreateUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

func GetUserHandler(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := biz.NewGetUserLogic(r.Context(), ctx)
		resp, err := l.GetUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

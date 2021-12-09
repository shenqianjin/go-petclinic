package handler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"petclinic/internal/biz"
	"petclinic/internal/model/dto"
	"petclinic/internal/support"
)

func CreateUserHandler(ctx *support.ServiceContext) http.HandlerFunc {
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

func GetUserHandler(ctx *support.ServiceContext) http.HandlerFunc {
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

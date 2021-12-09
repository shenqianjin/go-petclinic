package biz

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"petclinic/internal/model/dto"
	"petclinic/internal/service"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *service.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *service.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req dto.Request) (*dto.Response, error) {
	// todo: add your biz here and delete this line

	return &dto.Response{}, nil
}

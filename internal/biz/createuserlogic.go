package biz

import (
	"context"
	"petclinic/internal/model/dto"

	"github.com/tal-tech/go-zero/core/logx"
	"petclinic/internal/service"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *service.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *service.ServiceContext) CreateUserLogic {
	return CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req dto.Request) error {
	// todo: add your biz here and delete this line

	return nil
}

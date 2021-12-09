package biz

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"petclinic/internal/model/dto"
	"petclinic/internal/support"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *support.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *support.ServiceContext) CreateUserLogic {
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

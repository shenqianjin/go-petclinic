package service

import (
	"petclinic/internal/model"
)

type ServiceContext struct {
	Config model.Config
}

func NewServiceContext(c model.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

package service

import "petclinic/internal/support"

type ServiceContext struct {
	Config support.Config
}

func NewServiceContext(c support.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

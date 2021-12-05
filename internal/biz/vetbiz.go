package biz

import (
	"context"
	"log"
	"petclinic/internal/model"
	"petclinic/internal/service"
	"sync"
)

type petBiz struct {
	log    log.Logger
	ctx    context.Context
	svcCtx *service.ServiceContext
}

var (
	petBizOnce     sync.Once
	petBizInstance *petBiz
)

func GetPetBiz() *petBiz {
	petBizOnce.Do(func() {
		petBizInstance = &petBiz{}
	})
	return petBizInstance
}

func (b *petBiz) Vets() ([]*model.Vet, error) {
	// validate request
	// call service
	vets, err := service.GetVetService().FindVets(model.Vet{})
	// after validate
	return vets, err
}

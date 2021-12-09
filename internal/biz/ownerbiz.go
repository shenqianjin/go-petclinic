package biz

import (
	"context"
	"log"
	"petclinic/internal/model"
	"petclinic/internal/service"
	"petclinic/internal/support"
	"sync"
	"time"
)

type ownerBiz struct {
	log    log.Logger
	ctx    context.Context
	svcCtx *support.ServiceContext
}

var (
	ownerBizOnce     sync.Once
	ownerBizInstance *ownerBiz
)

// GetOwnerBiz get singleton owner biz object
//
// TODO: re-consider to rename it as NewXXX, prototype
//
func GetOwnerBiz() *ownerBiz {
	ownerBizOnce.Do(func() {
		ownerBizInstance = &ownerBiz{}
	})
	return ownerBizInstance
}

func (b *ownerBiz) Owners(lastName string) ([]*model.Owner, error) {
	// validate request
	// call service
	param := model.Owner{LastName: lastName}
	owners, err := service.GetOwnerService().FindOwners(param)
	// after validate
	return owners, err
}

func (b *ownerBiz) AddOwner(owner model.Owner) int64 {
	return service.GetOwnerService().SaveOwner(owner)
}

func (b *ownerBiz) FindOwner(id int64) (*model.Owner, error) {
	owner, err := service.GetOwnerService().FindOwner(id)
	visits := []model.Visit{
		{Date: time.Now().Add(time.Duration(-time.Hour * 72)), Description: "descrition 1"},
		{Date: time.Now().Add(time.Duration(-time.Hour * 360)), Description: "descrition 2"},
	}
	pets := []model.Pet{
		{Id: 1, Name: "Jack", BirthDate: time.Now().Add(time.Duration(-time.Hour * 7200)), Type: "dog", Visits: visits},
		{Id: 2, Name: "Mike", BirthDate: time.Now().Add(time.Duration(-time.Hour * 720)), Type: "cat", Visits: visits},
	}
	owner.Pets = pets
	return owner, err
}

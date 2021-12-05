package service

import (
	"petclinic/internal/dao"
	"petclinic/internal/dao/mem"
	"petclinic/internal/model"
	"sort"
	"sync"
)

type vetService struct {
	vetDao dao.VetDao
}

var (
	vetServiceOnce     sync.Once
	vetServiceInstance *vetService
)

func GetVetService() *vetService {
	vetServiceOnce.Do(func() {
		vetServiceInstance = &vetService{
			vetDao: mem.GetVetDao(),
		}
	})
	return vetServiceInstance
}

func (v *vetService) FindVets(param model.Vet) ([]*model.Vet, error) {
	// query data from Dao
	vets := v.vetDao.QueryList(param)
	// Sort by Id
	sort.Slice(vets, func(i, j int) bool {
		return vets[i].Id < vets[j].Id
	})
	// others
	return vets, nil
}

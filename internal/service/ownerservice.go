package service

import (
	"petclinic/internal/dao"
	"petclinic/internal/dao/mem"
	"petclinic/internal/model"
	"sort"
	"sync"
)

type ownerService struct {
	ownerDao dao.OwnerDao
}

var (
	ownerServiceOnce     sync.Once
	ownerServiceInstance *ownerService
)

func GetOwnerService() *ownerService {
	ownerServiceOnce.Do(func() {
		ownerServiceInstance = &ownerService{
			ownerDao: mem.GetOwnerDao(),
		}
	})
	return ownerServiceInstance
}

func (s *ownerService) FindOwners(param model.Owner) ([]*model.Owner, error) {
	// query data from Dao
	owners := s.ownerDao.QueryList(param)
	// Sort by Id
	sort.Slice(owners, func(i, j int) bool {
		return owners[i].Id < owners[j].Id
	})
	// others
	return owners, nil
}

func (s *ownerService) FindOwner(id int64) (*model.Owner, error) {
	// query data from Dao
	owner, err := s.ownerDao.Query(id)
	// others
	return owner, err
}

func (s *ownerService) SaveOwner(param model.Owner) int64 {
	result, _ := s.ownerDao.Insert(param)
	return result
}

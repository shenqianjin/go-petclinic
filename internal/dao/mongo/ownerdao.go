package mongo

import (
	"fmt"
	"petclinic/internal/model"
	"sync"
)

type ownerDaoImpl struct {
}

var (
	ownerDaoOnce sync.Once
	ownerDao     *ownerDaoImpl
)

func GetOwnerDao() *ownerDaoImpl {
	ownerDaoOnce.Do(func() {
		ownerDao = &ownerDaoImpl{}
	})
	return ownerDao
}

var ownerRepo = make(map[int64]*model.Owner)

func init() {
}

func (d *ownerDaoImpl) Insert(r model.Owner) (int64, error) {
	return 0, nil
}

func (d *ownerDaoImpl) Delete(id int64) *model.Owner {
	return nil
}

func (d *ownerDaoImpl) Upsert(r model.Owner) (int, error) {
	existed, ok := ownerRepo[r.Id]
	// auto generate Id
	if r.Id <= 0 {
		r.Id = int64(len(ownerRepo) + 1)
	}
	if ok {
		existed.FirstName = r.FirstName
		existed.LastName = r.LastName
		existed.City = r.City
		existed.Pets = r.Pets
		return 1, nil
	} else {
		ownerRepo[r.Id] = &r
		return 1, nil
	}
}

func (d *ownerDaoImpl) Query(id int64) (*model.Owner, error) {
	owner, ok := ownerRepo[id]
	if !ok {
		return nil, fmt.Errorf("owner[id = %v] isn't existed", id)
	}
	return owner, nil
}

func (d *ownerDaoImpl) QueryList(param model.Owner) []*model.Owner {
	result := make([]*model.Owner, 0)
	for _, v := range ownerRepo {
		if param.FirstName != "" && param.FirstName != v.FirstName {
			continue
		}
		if param.LastName != "" && param.LastName != v.LastName {
			continue
		}
		if param.City != "" && param.City != v.City {
			continue
		}
		result = append(result, v)
	}
	return result
}

package mem

import (
	"fmt"
	"petclinic/internal/model"
	"sync"
)

type petDaoImpl struct {
}

var (
	petDaoOnce sync.Once
	petDao *petDaoImpl
)

func GetPetDao() *petDaoImpl {
	petDaoOnce.Do(func() {
		petDao = &petDaoImpl{}
	})
	return petDao
}

var petRepo = make(map[int64]*model.Pet)

func (d *petDaoImpl) Insert(e model.Pet) (int, error) {
	if _, ok := petRepo[e.Id]; ok {
		return 0, fmt.Errorf("pet[id = %v] is already existed", e.Id)
	}
	// auto generate Id
	if e.Id <= 0 {
		e.Id = int64(len(petRepo) + 1)
	}
	petRepo[e.Id] = &e

	return 1, nil
}

func (d *petDaoImpl) Delete(id int64) *model.Pet {
	existed, ok := petRepo[id]
	if ok {
		delete(petRepo, id)
	}
	return existed
}

func (d *petDaoImpl) Upsert(e model.Pet) (int, error) {
	existed, ok := petRepo[e.Id]
	// auto generate Id
	if e.Id <= 0 {
		e.Id = int64(len(petRepo) + 1)
	}
	if ok {
		existed.Name = e.Name
		existed.BirthDate = e.BirthDate
		existed.Type = e.Type
		for _, v := range e.Visits {
			existed.Visits = append(existed.Visits, v)
		}
		for _, o := range e.Owner {
			existed.Owner = append(existed.Owner, o)
		}
		return 1, nil
	} else {
		petRepo[e.Id] = &e
		return 1, nil
	}
}

func (d *petDaoImpl) Query(id int64) *model.Pet {
	return petRepo[id]
}

func (d *petDaoImpl) QueryList(param model.Pet) []*model.Pet {
	result := make([]*model.Pet, 0)
	for _, v := range petRepo {
		if param.Name != "" && param.Name != v.Name {
			continue
		}
		if param.Type != "" && param.Type != v.Type {
			continue
		}
		result = append(result, v)
	}
	return result
}


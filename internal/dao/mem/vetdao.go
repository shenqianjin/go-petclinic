package mem

import (
	"fmt"
	"os"
	"petclinic/internal/model"
	"sync"
)

type vetDaoImpl struct {
}

var (
	vetDaoOnce sync.Once
	vetDao *vetDaoImpl
)

func GetVetDao() *vetDaoImpl {
	vetDaoOnce.Do(func() {
		vetDao = &vetDaoImpl{}
	})
	return vetDao
}

var vetRepo = make(map[int64]*model.Vet)

func init() {
	dbType := os.Getenv("DB_TYPE")
	if dbType != "" && dbType != "MEM" {
		return
	}
	vetRepo[1] = &model.Vet{Id: 1, FirstName: "James", LastName: "Carter", City: "NewYork"}
	vetRepo[2] = &model.Vet{Id: 2, FirstName: "Helen", LastName: "Leary", City: "HongKong", Specialties: []string{"radiology"}}
	vetRepo[3] = &model.Vet{Id: 3, FirstName: "Linda", LastName: "Douglas", City: "London", Specialties: []string{"dentistry", "surgery"}}
	vetRepo[4] = &model.Vet{Id: 4, FirstName: "Rafael", LastName: "Ortega", City: "Shanghai", Specialties: []string{"surgery"}}
	vetRepo[5] = &model.Vet{Id: 5, FirstName: "Henry", LastName: "Stevens", City: "Beijing", Specialties: []string{"radiology"}}
	vetRepo[6] = &model.Vet{Id: 6, FirstName: "Sharon", LastName: "Jenkins", City: "NewYork"}
}

func (d *vetDaoImpl) Insert(r model.Vet) (int, error) {
	if _, ok := vetRepo[r.Id]; ok {
		return 0, fmt.Errorf("vet[id = %v] is already existed", r.Id)
	}
	// auto generate Id
	if r.Id <= 0 {
		r.Id = int64(len(vetRepo) + 1)
	}
	vetRepo[r.Id] = &r

	return 1, nil
}

func (d *vetDaoImpl) Delete(id int64) *model.Vet {
	existed, ok := vetRepo[id]
	if ok {
		delete(vetRepo, id)
	}
	return existed
}

func (d *vetDaoImpl) Upsert(r model.Vet) (int, error) {
	existed, ok := vetRepo[r.Id]
	// auto generate Id
	if r.Id <= 0 {
		r.Id = int64(len(vetRepo) + 1)
	}
	if ok {
		existed.FirstName = r.FirstName
		existed.LastName = r.LastName
		existed.City = r.City
		existed.Specialties = r.Specialties
		return 1, nil
	} else {
		vetRepo[r.Id] = &r
		return 1, nil
	}
}

func (d *vetDaoImpl) Query(id int64) *model.Vet {
	return vetRepo[id]
}

func (d *vetDaoImpl) QueryList(param model.Vet) []*model.Vet {
	result := make([]*model.Vet, 0)
	for _, v := range vetRepo {
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


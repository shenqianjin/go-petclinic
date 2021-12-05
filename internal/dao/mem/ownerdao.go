package mem

import (
	"fmt"
	"os"
	"petclinic/internal/model"
	"sync"
)

type ownerDaoImpl struct {
}

var (
	ownerDaoOnce sync.Once
	ownerDao *ownerDaoImpl
)

func GetOwnerDao() *ownerDaoImpl {
	ownerDaoOnce.Do(func() {
		ownerDao = &ownerDaoImpl{}
	})
	return ownerDao
}

var ownerRepo = make(map[int64]*model.Owner)

func init() {
	dbType := os.Getenv("DB_TYPE")
	if dbType != "" && dbType != "MEM" {
		return
	}

	ownerRepo[1] = &model.Owner{Id: 1, FirstName: "George", LastName: "Franklin", Address: "110 W. Liberty St.", City: "NewYork", Telephone: "6085551023", Pets: []model.Pet{{Name: "Leo"}}}
	ownerRepo[2] = &model.Owner{Id: 2, FirstName: "Betty", LastName: "Davis", Address: "638 Cardinal Ave.", City: "Sun Prairie", Telephone: "6085551749", Pets: []model.Pet{{Name: "Basil"}}}
	ownerRepo[3] = &model.Owner{Id: 3, FirstName: "Eduardo", LastName: "Rodriquez", Address: "2693 Commerce St.", City: "McFarland", Telephone: "6085558763", Pets: []model.Pet{{Name: "Jewel"}, {Name: "Rosy"}}}
	ownerRepo[4] = &model.Owner{Id: 4, FirstName: "Harold", LastName: "Davis", Address: "563 Friendly St.", City: "Windsor", Telephone: "6085553198", Pets: []model.Pet{{Name: "Iggy"}}}
	ownerRepo[5] = &model.Owner{Id: 5, FirstName: "Peter", LastName: "McTavish", Address: "2387 S. Fair Way", City: "Madison", Telephone: "6085552765", Pets: []model.Pet{{Name: "George"}}}

	ownerRepo[6] = &model.Owner{Id: 6, FirstName: "Jean", LastName: "Coleman", Address: "105 N. Lake St.", City: "Monona", Telephone: "6085552654", Pets: []model.Pet{{Name: "Max"}, {Name: "Samantha"}}}
	ownerRepo[7] = &model.Owner{Id: 7, FirstName: "Jeff", LastName: "Black", Address: "1450 Oak Blvd.", City: "Monona", Telephone: "6085555387", Pets: []model.Pet{{Name: "Lucky"}}}
	ownerRepo[8] = &model.Owner{Id: 8, FirstName: "Maria", LastName: "Escobito", Address: "345 Maple St.", City: "Madison", Telephone: "6085557683", Pets: []model.Pet{{Name: "Mulligan"}}}
	ownerRepo[9] = &model.Owner{Id: 9, FirstName: "David", LastName: "Schroeder", Address: "2749 Blackhawk Trail", City: "Madison", Telephone: "6085559435", Pets: []model.Pet{{Name: "Freddy"}}}
	ownerRepo[10] = &model.Owner{Id: 10, FirstName: "Carlos", LastName: "Estaban", Address: "2335 Independence La.", City: "Waunakee", Telephone: "6085555487", Pets: []model.Pet{{Name: "Lucky"}, {Name: "Sly"}}}
}

func (d *ownerDaoImpl) Insert(r model.Owner) (int64, error) {
	if _, ok := ownerRepo[r.Id]; ok {
		return 0, fmt.Errorf("owner[id = %v] is already existed", r.Id)
	}
	// auto generate Id
	if r.Id <= 0 {
		r.Id = int64(len(ownerRepo) + 1)
	}
	ownerRepo[r.Id] = &r

	return r.Id, nil
}

func (d *ownerDaoImpl) Delete(id int64) *model.Owner {
	existed, ok := ownerRepo[id]
	if ok {
		delete(ownerRepo, id)
	}
	return existed
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


package dao

import "petclinic/internal/model"

type PetDao interface {
	Insert(r model.Pet) (int, error)
	Delete(id int64) *model.Pet
	Upsert(r model.Pet) (int, error)
	Query(id int64) *model.Pet
	QueryList(param model.Pet) []*model.Pet
}

type VetDao interface {
	Insert(r model.Vet) (int, error)
	Delete(id int64) *model.Vet
	Upsert(r model.Vet) (int, error)
	Query(id int64) *model.Vet
	QueryList(param model.Vet) []*model.Vet
}

type OwnerDao interface {
	Insert(r model.Owner) (int64, error)
	Delete(id int64) *model.Owner
	Upsert(r model.Owner) (int, error)
	Query(id int64) (*model.Owner, error)
	QueryList(param model.Owner) []*model.Owner
}

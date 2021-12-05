package model

import "time"

type NameEntity struct {
	Name string
}

type Person struct {
	FirstName   string
	LastName    string
}

type Vet struct {
	Id int64
	FirstName   string
	LastName    string
	City        string
	Specialties []string
}

type PetType string

type Pet struct {
	Id int64
	Name string
	BirthDate time.Time
	Type string
	Owner []Owner
	Visits []Visit

}

type Visit struct {
	Date time.Time
	Description string
	PetId int64

}

type Owner struct {
	Id int64
	FirstName   string
	LastName    string
	Address string
	City string
	Telephone string
	Pets []Pet
}



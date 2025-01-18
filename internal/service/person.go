package service

import (
	"testing/internal/model"
)

type PersonRepo interface {
	Get(name string) (model.Person, error)
}

type PersonService struct {
	repo PersonRepo
}

func NewPersonService(personRepo PersonRepo) *PersonService {
	return &PersonService{
		repo: personRepo,
	}
}

func (p *PersonService) Get(name string) (model.Person, error) {
	return p.repo.Get(name)
}

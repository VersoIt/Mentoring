package repository

import (
	"log"
	"testing/internal/err"
	"testing/internal/model"
)

type PersonRepository struct {
	personMap map[string]model.Person
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{
		personMap: map[string]model.Person{
			"Ruslan": {
				Name:   "Ruslan",
				Age:    20,
				Weight: 78,
				Height: 183,
			},
			"Artem": {
				Name:   "Artem",
				Age:    15,
				Weight: 72,
				Height: 177,
			},
		},
	}
}

func (p *PersonRepository) Get(name string) (model.Person, error) {
	log.Printf("getting person by name %s...\n", name)

	person, ok := p.personMap[name]
	if !ok {
		log.Printf("person with name %s not found\n", name)
		return model.Person{}, err.NotFound
	}

	log.Printf("person with name %s successfully found!\n", name)

	return person, nil
}

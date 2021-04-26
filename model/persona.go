package model

import "github.com/jaswdr/faker"

type Persona struct {
	Id       int
	Nombre   string
	Apellido string
}

func NewPersona(id int, nombre string, apellido string) *Persona {
	return &Persona{
		Id:       id,
		Nombre:   nombre,
		Apellido: apellido,
	}
}
func Initialize() []Persona {
	var personas []Persona
	faker := faker.New()
	p := faker.Person()
	for i := 0; i < 100; i++ {
		var nuevaPersona = NewPersona(i, p.FirstName(), p.LastName())
		personas = append(personas, *nuevaPersona)
	}
	return personas
}

var PersonasDb = Initialize()

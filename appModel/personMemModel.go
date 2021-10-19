package appModel

import "fmt"

type PersonMemModel struct {
	data []Person
}

func NewPersonMemModel() *PersonMemModel {
	return &PersonMemModel{
		data: []Person{},
	}
}

func (pm *PersonMemModel) GetByEmailAndPassword(email string, password string) (Person, error) {
	var person = Person{}
	for _, person = range pm.data {
		if person.Email == email && person.Password == password {
			return person, nil
		}
	}
	return person, fmt.Errorf("person not found")
}

func (pm *PersonMemModel) GetAll() ([]Person, error) {
	return pm.data, nil
}

func (pm *PersonMemModel) Add(p Person) (Person, error) {
	p.ID = uint(len(pm.data))
	pm.data = append(pm.data, p)
	return p, nil
}

func (pm *PersonMemModel) Edit(id int, p Person) (Person, error) {
	if id < 0 || id >= len(pm.data) {
		return p, fmt.Errorf("person %d not found", id)
	}
	p.ID = uint(id)
	pm.data[id] = p
	return p, nil
}

package appModel

type PersonModel interface {
	GetByEmailAndPassword(email string, password string) (Person, error)
	GetAll() ([]Person, error)
	Add(Person) (Person, error)
	Edit(int, Person) (Person, error)
}

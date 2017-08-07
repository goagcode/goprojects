package databaselayer

const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

type DinoDBHandler interface {
	GetAvailableDynos() ([]Animal, error)
	GetDynoByNickanem() (Animal, error)
	GetDynosByType() ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

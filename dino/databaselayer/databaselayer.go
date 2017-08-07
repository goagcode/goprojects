package databaselayer

import "errors"

const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

type DinoDBHandler interface {
	GetAvailableDynos() ([]Animal, error)
	GetDynoByNickanem(string) (Animal, error)
	GetDynosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

var DBTypeNotSupported = errors.New("The Database type provided is not supported...")

func GetDatabaseHandler(dbtype uint8, connection string) (DinoDBHandler, error) {
	switch dbtype {
	case MONGODB:
		return NewMongoDBHandler(connection)
	}
	return nil, DBTypeNotSupported
}

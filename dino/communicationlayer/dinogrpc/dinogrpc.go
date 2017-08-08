package dinogrpc

import (
	"fmt"

	context "golang.org/x/net/context"

	"github.com/miguellgt/goprojects/dino/databaselayer"
)

type DinoGrpcServer struct {
	dbHandler databaselayer.DinoDBHandler
}

func NewDinoGrpcServer(dbtype uint8, connstring string) (*DinoGrpcServer, error) {
	handler, err := databaselayer.GetDataBaseHandler(dbtype, connstring)
	if err != nil {
		return nil, fmt.Errorf("Could not create a database handler object, error %v", err)
	}
	return &DinoGrpcServer{
		dbHandler: handler,
	}, nil
}

func (server *DinoGrpcServer) GetAnimal(ctx context.Context, r *Request) (*Animal, error) {
	animal, err := server.dbHandler.GetDynoByNickname(r.GetNickname())
	return convertToDinoGRPCAnimal(animal), err
}

func convertToDinoGRPCAnimal(animal databaselayer.Animal) *Animal {
	return &Animal{
		Id:         int32(animal.ID),
		AnimalType: animal.AnimalType,
		Nickname:   animal.Nickname,
		Zone:       int32(animal.Zone),
		Age:        int32(animal.Age),
	}
}

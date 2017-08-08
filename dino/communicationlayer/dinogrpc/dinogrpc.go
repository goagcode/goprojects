package dinogrpc

import (
	"fmt"

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
	}
}

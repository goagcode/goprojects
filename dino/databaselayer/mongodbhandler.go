package databaselayer

import mgo "gopkg.in/mgo.v2"

type MongoDBHandler struct {
	*mgo.Session
}

func NewMongoDBHandler(connection string) (*MongoDBHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongoDBHandler{
		Session: s,
	}, err
}

func (handler *MongoDBHandler) GetAvailableDynos() ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("dino").C("animals").Find(nil).All(&animals)
	return animals, err
}

func (handler *MongoDBHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}

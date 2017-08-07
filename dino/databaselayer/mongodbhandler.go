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

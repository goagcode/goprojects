package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Animal struct {
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	// Get a Collection
	animalCollection := session.DB("dino").C("animals")

	animals := []interface{}{
		Animal{
			AnimalType: "Tynanousario rex",
			Nickname:   "rex",
			Zone:       1,
			Age:        14,
		},
		Animal{
			AnimalType: "Velociraptor",
			Nickname:   "rapto",
			Zone:       2,
			Age:        17,
		},
	}
	err = animalCollection.Insert(animals...)
	if err != nil {
		log.Fatal(err)
	}
	// Update Document
	err = animalCollection.Update(bson.M{"nickname": "rex"}, bson.M{"$set": bson.M{"age": 10}})
	if err != nil {
		log.Fatal(err)
	}

	// Remove Document
	err = animalCollection.Remove(bson.M{"nickname": "rex"})
	if err != nil {
		log.Fatal(err)
	}

	// Query
	// age > 10 and zone in (1, 2)
	query := bson.M{
		"age": bson.M{
			"$gt": 10,
		},
		"zone": bson.M{
			"$in": []int{1, 2},
		},
	}
	results := []Animal{}
	animalCollection.Find(query).All(&results) //.one
	fmt.Println(results)
}

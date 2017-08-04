package main

import (
	"log"

	"github.com/influxdata/influxdb/client/v2"
)

// Collect the Weights of each animal on a frequent basis => time series dataset
// this data get stored in influxdb, so that we can use it later

var animalTags = []string{
	"Tynanousario rex;Rex",
	"Velociraptor;Raptor",
	"Velociraptor;Velo",
	"Carnotaurus;Carno",
}

const myDB = "dino"

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	queryDB(c, "", "CREATE DATABASE"+myDB)
	// Create a batch point object
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  myDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func queryDB(c client.Client, database, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: Database,
	}
	response, err := c.Query(q)
	if err != nil {
		return res, err
	}
	if response.Error() != nil {
		return res, response.Error()
	}
	return response.Results, nil
}

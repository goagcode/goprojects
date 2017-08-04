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
}

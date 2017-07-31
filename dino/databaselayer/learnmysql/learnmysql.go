package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	Id         int
	AnimalType string
	Nickname   string
	Zone       int
	Age        int
}

func main() {
	// connect to the database
	db, err := sql.Open("mysql", "root:root@/dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Coonected to database")
}

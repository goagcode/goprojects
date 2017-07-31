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
	// General query with argumets
	rows, err := db.Query("SELECT * FROM animals")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.Id, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	fmt.Println(animals)
	result, err := db.Exec("INSERT INTO animals(animal_type, nickname, zone, age) VALUES('Carnotaurus', ?, 2, 28)", "carnito")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())
	age := 14
	id := 2
	result, err := db.Exec("UPDATE animals SET age = ? where id = ?", age, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}

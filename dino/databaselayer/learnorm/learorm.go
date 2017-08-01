package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	//gorm.Model
	ID         int    `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	AnimalType string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT"`
	Zone       int    `gorm:"type:INTEGER"`
	Age        int    `gorm:"type:INTEGER"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/dino?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("It is ok")
}

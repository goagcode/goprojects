package main

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	// gorm.Model
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

	db.Table("animals").DropTableIfExists(&Animal{})
	db.Table("dinos").DropTableIfExists(&Animal{})
	db.Table("animals").AutoMigrate(&Animal{})
	db.Table("dinos").AutoMigrate(&Animal{})

	a := Animal{
		AnimalType: "Tynanousario Rex",
		Nickname:   "rex",
		Zone:       1,
		Age:        11,
	}

	db.Create(&a)
	db.Table("dinos").Create(&a)

	a = Animal{
		AnimalType: "Velociraptor",
		Nickname:   "rapto",
		Zone:       2,
		Age:        15,
	}

	db.Save(&a)
	// Update animal table filds
	db.Table("animals").Where("nickname = ? AND zone = ?", "rex", 1).Update("age", 12)
}

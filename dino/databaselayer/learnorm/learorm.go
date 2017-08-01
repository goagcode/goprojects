package main

import (
	"fmt"
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

	db.Debug().Save(&a)

	// Update animal table filds
	// UPDATE ANIMALS SET AGE = 12 WHERE NICKNAME = "rex" AND ZONE = 1
	db.Debug().Table("animals").Where("nickname = ? AND zone = ?", "rex", 1).Update("age", 12)

	// Query
	// SELECT * FROM ANIMALS WHERE AGE > 12
	animals := []Animal{}
	db.Debug().Find(&animals, "age > ?", 12)
	fmt.Println(animals)
}

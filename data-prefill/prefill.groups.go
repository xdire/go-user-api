package data_prefill

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xdire/go-user-api/models"
)

func GenerateInitialGroupSet(db *gorm.DB)  {

	db.Create(&models.Group{
		Name: "Admin",
	})

	db.Create(&models.Group{
		Name: "QA",
	})

	db.Create(&models.Group{
		Name: "Sales",
	})

	db.Create(&models.Group{
		Name: "Engineering",
	})

	db.Create(&models.Group{
		Name: "Authorized",
	})

}

func TestInitialGroupSet(db *gorm.DB)  {

	var groupOne models.Group
	db.First(&groupOne, 1)
	fmt.Printf("Group one found: %v \n", groupOne)

	var groupEngineering models.Group
	db.First(&groupEngineering, "name = ?", "Engineering")
	fmt.Printf("Group engineering found: %v \n", groupEngineering)

}

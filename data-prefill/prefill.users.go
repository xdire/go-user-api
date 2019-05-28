package data_prefill

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xdire/go-user-api/models"
)

func GenerateInitialUserSet(db *gorm.DB)  {

	err := db.Create(&models.User{
		Email: "test@test.com",
		Login: "test1",
		FirstName: "John",
		LastName: "Wick"}).Error
	if err != nil {
		fmt.Println("User John Wick already created")
	}

	err = db.Create(&models.User{
		Email: "sales@test.com",
		Login: "sales1",
		FirstName: "Mary",
		LastName: "Finkilshtein"}).Error
	if err != nil {
		fmt.Println("User Mary Finkilshtein already created")
	}

	err = db.Create(&models.User{
		Email: "support@test.com",
		Login: "support1",
		FirstName: "Alex",
		LastName: "Pouzo"}).Error
	if err != nil {
		fmt.Println("User Alex Pouzo already created")
	}

	err = db.Create(&models.User{
		Email: "eng@test.com",
		Login: "engineer1",
		FirstName: "Peter",
		LastName: "Luri"}).Error
	if err != nil {
		fmt.Println("User Peter Luri already created")
	}

	err = db.Create(&models.User{
		Email: "qa@test.com",
		Login: "horror1",
		FirstName: "Liu",
		LastName: "Kang"}).Error
	if err != nil {
		fmt.Println("User Liu Kang already created")
	}

}

func TestInitialUserSet(db *gorm.DB)  {

	var userOne models.User
	db.First(&userOne, 1)
	fmt.Printf("User one found: %v \n", userOne)

	var userSales models.User
	db.First(&userSales, "login = ?", "sales1")
	fmt.Printf("User sales found: %v \n", userSales)

	var userEmailEng models.User
	db.First(&userEmailEng, "email = ?", "eng@test.com")
	fmt.Printf("User email found: %v \n", userEmailEng)

}
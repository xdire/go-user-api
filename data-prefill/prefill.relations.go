package data_prefill

import (
	"github.com/jinzhu/gorm"
	"github.com/xdire/go-user-api/models"
)

func GenerateUserGroupSalesRelations(db *gorm.DB)  {

	var userSales models.User
	db.First(&userSales, "login = ?", "sales1")

	var groupSales models.Group
	db.First(&groupSales, "name = ?", "Sales")

	db.Model(userSales).Association("Groups").Append(groupSales)

}

func GenerateUserGroupAuthorizedRelations(db *gorm.DB)  {

	var users []models.User
	db.Where("login != ?", "test1").Find(&users)

	var groupAuthorized models.Group
	db.First(&groupAuthorized, "name = ?", "Authorized")

	for _, user := range(users) {
		db.Model(user).Association("Groups").Append(groupAuthorized)
	}

}

package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uint64 	`gorm:"primary_key"`
	Uid       string 	`gorm:"unique_index:idx_user_uid"`
	Login     string 	`gorm:"unique_index:idx_login"`
	Email     string 	`gorm:"unique_index:idx_email"`
	FirstName string
	LastName  string
	Groups    []*Group	`gorm:"many2many:user_groups"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this *User) BeforeSave() error {
	if len(this.Uid) == 0 {
		uid, err := uuid.NewRandom()
		if err != nil { return err }
		this.Uid = uid.String()
	}
	return nil
}

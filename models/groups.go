package models

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ID        uint64  `gorm:"primary_key"`
	Uid       string  `gorm:"unique_index:idx_group_uid"`
	Name      string  `gorm:"unique_index:idx_group_name"`
	Users     []*User `gorm:"many2many:user_groups"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this *Group) BeforeSave() error {
	if len(this.Uid) == 0 {
		uid, err := uuid.NewRandom()
		if err != nil { return err }
		this.Uid = uid.String()
	}
	return nil
}

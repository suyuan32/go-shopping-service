//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName      string    `gorm:"type:varchar(50);not null"`
	LoginName     string    `gorm:"type:varchar(50);not null;"`
	Email         string    `gorm:"type:varchar(50);not null;index:idx_mail,unique;"`
	Mobile        string    `gorm:"type:varchar(50);not null;"`
	LoginPassword string    `gorm:"type:varchar(255);not null;"`
	Pic           string    `gorm:"type:varchar(255);not null;"`
	Status        int32     `gorm:"type:tinyint;not null;default:1;comment:1 valid 0 invalid"`
	LastLoginAt   time.Time `gorm:"type:datetime;"`
	LastLoginIP   string    `gorm:"type:varchar(20);default:0.0.0.0;not null;"`
	RoleID        uint
	Role          Role
}

type Role struct {
	gorm.Model
	RoleName string `gorm:"type:varchar(20);not null"`
	Remark   string `gorm:"type:varchar(500);default:'';not null"`
}

type RoleMenu struct {
	gorm.Model
	RoleID uint
	Role   Role
	MenuID uint
	Menu   Menu
}

type Menu struct {
	gorm.Model
	ParentID   uint
	Parent     *Menu
	Name       string `gorm:"type:varchar(20);not null"`
	URL        string `gorm:"type:varchar(255);not null"`
	Permission string `gorm:"type:varchar(255);not null"`
	Type       int8   `gorm:"type:tinyint;default:0;not null"`
	Icon       string `gorm:"type:varchar(255);not null"`
	OrderNum   int32  `gorm:"type:int;default:999;not null"`
}

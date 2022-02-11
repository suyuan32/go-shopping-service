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
	RealName      string    `gorm:"type:varchar(50);not null;"`
	Email         string    `gorm:"type:varchar(50);not null;index:idx_mail,unique;"`
	LoginPassword string    `gorm:"type:varchar(255);not null;"`
	PayPassword   string    `gorm:"type:varchar(50);not null;"`
	Mobile        string    `gorm:"type:varchar(50);not null;"`
	Sex           string    `gorm:"type:char(1);default:M;not null;"`
	Birthday      time.Time `gorm:"type:datetime;"`
	Pic           string    `gorm:"type:varchar(255);not null;"`
	Status        int8      `gorm:"type:tinyint;not null;default:1;comment:1 valid 0 invalid"`
	LastLoginAt   time.Time `gorm:"type:datetime;"`
	LastLoginIP   string    `gorm:"type:varchar(20);default:0.0.0.0;not null;"`
	Memo          string    `gorm:"type:varchar(500);default:'';not null;"`
	Score         int32     `gorm:"type:int;default:0;not null;comment:user score is used to get discount."`
}

type UserAddress struct {
	gorm.Model
	UserId          uint   `gorm:"type:bigint(20);not null"`
	Receiver        string `gorm:"type:varchar(20);not null"`
	ProvinceId      uint   `gorm:"type:int;not null"`
	Province        string `gorm:"type:varchar(20);not null"`
	CityId          uint   `gorm:"type:int;not null"`
	City            string `gorm:"type:varchar(20);not null"`
	AreaId          uint   `gorm:"type:int;not null"`
	Area            string `gorm:"type:varchar(20);not null"`
	PostCode        string `gorm:"type:varchar(15);not null"`
	Address         string `gorm:"type:varchar(100);not null"`
	Mobile          string `gorm:"type:varchar(20);not null"`
	Status          int8   `gorm:"type:tinyint;default:0;not null"`
	IsCommonAddress bool   `gorm:"type:bool;default:false;not null"`
}

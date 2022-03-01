//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package model

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	UserID    uint32 `gorm:"type:bigint(20);not null"`
	Operation string `gorm:"type:varchar(300);not null"`
	Method    string `gorm:"type:varchar(50);not null"`
	IP        string `gorm:"type:varchar(20);not null"`
}

//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package handler

import (
	"gorm.io/gorm"
	"sys_srv/proto"
)

type SystemServer struct {
	proto.UnimplementedSystemServer
}

func Paginate(page, pageSize uint32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

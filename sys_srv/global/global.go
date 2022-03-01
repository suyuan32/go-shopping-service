//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package global

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
	"sys_srv/config"
)

var (
	NacosConfig  = &config.NacosConfig{}
	ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
	Lang         *i18n.Localizer
)

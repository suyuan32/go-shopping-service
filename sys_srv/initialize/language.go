//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"

	"sys_srv/global"
)

func InitI18n() {
	var langTag language.Tag
	var filePath string
	if global.ServerConfig.Lang == "zh" {
		langTag = language.Chinese
		filePath = fmt.Sprintf("%ssys_srv/language/active.zh.toml", viper.GetString("GO_SHOPPING_SRV"))
	} else if global.ServerConfig.Lang == "en" {
		langTag = language.English
		filePath = fmt.Sprintf("%ssys_srv/language/active.en.toml", viper.GetString("GO_SHOPPING_SRV"))
	}
	bundle := i18n.NewBundle(langTag)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(filePath)

	global.Lang = i18n.NewLocalizer(bundle, global.ServerConfig.Lang)
}

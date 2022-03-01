//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import (
	"go.uber.org/zap"

	"sys_srv/global"
)

func InitZapLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(global.GetMessage("FailStartZap"))
	}
	zap.ReplaceGlobals(logger)
}

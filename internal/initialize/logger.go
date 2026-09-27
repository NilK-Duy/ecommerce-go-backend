package initialize

import (
	"ecommerce-backend/global"
	"ecommerce-backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}

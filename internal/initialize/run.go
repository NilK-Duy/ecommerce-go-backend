package initialize

import (
	"ecommerce-backend/global"
	"fmt"
	"log"

	"go.uber.org/zap"
)

func Run() {
	// load configuration
	LoadingConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "Success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

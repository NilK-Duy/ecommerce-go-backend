package global

import (
	"ecommerce-backend/pkg/logger"
	"ecommerce-backend/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)

/*
Config
Redis
Mysql
...
*/

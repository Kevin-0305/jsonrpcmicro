package global

import (
	"jsonrpcmicro/api/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	REDIS     *redis.Client
	Log       *logrus.Logger
	ApiConfig *config.Config
)

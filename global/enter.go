package global

import (
	"fast_gin/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const Version = "0.0.2"

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)

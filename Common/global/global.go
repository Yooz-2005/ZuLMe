package global

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
	Es  *elasticsearch.Client
)

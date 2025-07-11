package global

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/minio/minio-go/v7"

	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	Rdb         *redis.Client
	Es          *elasticsearch.Client
	MongoClient *mongo.Client
	Ctx         context.Context
	Cancel      context.CancelFunc
	Minio       *minio.Client
)

package initialize

import (
	"ZuLMe/ZuLMe/Common/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var (
	dbOnce sync.Once
	dbErr  error
)

// MysqlInit 初始化 MySQL 连接并返回数据库实例
func MysqlInit() (*gorm.DB, error) {
	dbOnce.Do(func() {
		data := Nacos.Mysql
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			data.User, data.Password, data.Host, data.Port, data.Database)

		global.DB, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if dbErr != nil {
			log.Printf("mysql连接失败: %v", dbErr)
			return
		}

		sqlDB, err := global.DB.DB()
		if err != nil {
			dbErr = err
			log.Printf("获取底层数据库连接失败: %v", err)
			return
		}

		// 配置连接池
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)

		fmt.Println("mysql连接成功")
	})

	// 返回错误
	if dbErr != nil {
		return nil, dbErr
	}
	return global.DB, nil
}

package initialize

import (
	"Common/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit() {
	var err error
	data := Nacos.Mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", data.User, data.Password, data.Host, data.Port, data.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("mysql failed to connect")
	}
	fmt.Println("mysql connect success")
}

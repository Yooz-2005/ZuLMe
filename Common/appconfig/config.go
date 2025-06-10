package appconfig

import (
	"fmt"
	"github.com/spf13/viper"
)

type ViperConfData struct {
	Nacos struct {
		SpaceId string
		Host    string
		Port    uint64
		DataId  string
		GroupId string
	}
	SendSms struct {
		Count int64
	}
}

var ConfData ViperConfData

func GetViperConfData() {
	viper.SetConfigFile("../../Common/appconfig/config.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&ConfData)
	fmt.Println(ConfData)
	fmt.Println("viper is ok")
}

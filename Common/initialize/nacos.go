package initialize

import (
	"Common/appconfig"
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Nacoss struct {
	Mysql struct {
		User     string `json:"User"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		Database string `json:"Database"`
	} `json:"mysql"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Elasticsearch struct {
		Host string `json:"host"`
	} `json:"elasticsearch"`
	Group struct {
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"group"`
	Consul struct {
		Name string `json:"Name"`
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"consul"`
	Rabbitmq struct {
		Address string `json:"address"`
	} `json:"rabbitmq"`
	ALiYun struct {
		AccessKeyId     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
		Address         string `json:"address"`
		Bucket          string `json:"bucket"`
	} `json:"ALiYun"`
	ALiPay struct {
		PrivateKey string `json:"privateKey"`
		AppId      string `json:"appId"`
		NotifyUrl  string `json:"notifyUrl"`
		ReturnUrl  string `json:"returnUrl"`
	} `json:"ALiPay"`
	TenXunYun struct {
		SecretID  string `json:"SecretID"`
		SecretKey string `json:"SecretKey"`
	} `json:"tenXunYun"`
	QqEmail struct {
		Qq       string `json:"qq"`
		Password string `json:"password"`
	} `json:"qqEmail"`
	Minio struct {
		AccessKeyId     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
		Location        string `json:"location"`
		Bucket          string `json:"bucket"`
	} `json:"minio"`
	Mongodb struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		Host         string `json:"host"`
		Port         int    `json:"port"`
		DatabaseName string `json:"databaseName"`
	} `json:"mongodb"`
}

var Nacos Nacoss

func NewNacos() {
	cos := appconfig.ConfData.Nacos
	clientConfig := constant.ClientConfig{
		NamespaceId:         cos.SpaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      cos.Host,
			ContextPath: "/nacos",
			Port:        cos.Port,
			Scheme:      "http",
		},
	}

	client, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, _ := client.GetConfig(vo.ConfigParam{
		DataId: cos.DataId,
		Group:  cos.GroupId})
	json.Unmarshal([]byte(content), &Nacos)

	err := client.ListenConfig(vo.ConfigParam{
		DataId: cos.DataId,
		Group:  cos.GroupId,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + cos.GroupId + ", dataId:" + cos.DataId + ", data:" + data)
			json.Unmarshal([]byte(data), &Nacos)
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}

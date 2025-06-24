package initialize

import (
	"Common/global"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

// InitES 初始化 Elasticsearch 客户端
func InitES() {
	var err error

	// 从Nacos配置获取ES地址
	esHost := Nacos.Elasticsearch.Host
	if esHost == "" {
		esHost = "http://localhost:9200" // 默认地址
	} else {
		// 确保地址包含协议前缀
		if !strings.HasPrefix(esHost, "http://") && !strings.HasPrefix(esHost, "https://") {
			esHost = "http://" + esHost
		}
		// 如果端口是5601（Kibana），改为9200（ES）
		if strings.Contains(esHost, ":5601") {
			esHost = strings.Replace(esHost, ":5601", ":9200", 1)
		}
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			esHost,
		},
		// 添加一些基本配置
		Username: "", // 如果需要认证
		Password: "", // 如果需要认证
	}

	global.Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch cannot connect: %v\n", err)
		return
	}

	// 测试连接
	res, err := global.Es.Info()
	if err != nil {
		fmt.Printf("elasticsearch connection test failed: %v\n", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Printf("elasticsearch connection error: %s\n", res.String())
		return
	}

	fmt.Printf("elasticsearch connect success to: %s\n", esHost)
}

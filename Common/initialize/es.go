package initialize

import (
	"ZuLMe/ZuLMe/Common/global"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
)

// InitES 初始化 Elasticsearch 客户端
func InitES() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
		// ...
	}
	global.Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("elasticsearch cannot connect")
	}
	fmt.Println("elasticsearch connect success")

}

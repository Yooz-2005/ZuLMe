package main

import (
	"Common/appconfig"
	"Common/initialize"
	"Common/utils"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// 定义命令行参数
	var action = flag.String("action", "init", "操作类型: init(初始化数据) 或 clean(清理数据)")
	flag.Parse()

	fmt.Println("=== ZuLMe 数据初始化工具 ===")
	fmt.Printf("操作类型: %s\n", *action)

	// 初始化配置和数据库连接
	fmt.Println("正在初始化配置...")
	appconfig.GetViperConfData()
	
	fmt.Println("正在连接数据库...")
	initialize.MysqlInit()
	
	fmt.Println("正在连接Redis...")
	initialize.RedisInit()

	// 根据参数执行不同操作
	switch *action {
	case "init":
		fmt.Println("\n开始初始化商家测试数据...")
		err := utils.InitMerchantTestData()
		if err != nil {
			log.Fatalf("初始化数据失败: %v", err)
		}
		fmt.Println("✅ 数据初始化完成！")
		
	case "clean":
		fmt.Println("\n开始清理商家测试数据...")
		err := utils.CleanMerchantTestData()
		if err != nil {
			log.Fatalf("清理数据失败: %v", err)
		}
		fmt.Println("✅ 数据清理完成！")
		
	default:
		fmt.Printf("❌ 未知操作类型: %s\n", *action)
		fmt.Println("支持的操作类型:")
		fmt.Println("  init  - 初始化测试数据")
		fmt.Println("  clean - 清理测试数据")
		os.Exit(1)
	}
}

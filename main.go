package main

import (
	"flag"
	"fmt"
	"go_pushfile/file_rhash"
)

func p1() {
	fmt.Printf("初始化中\n")
	file_rhash.Json_writ()
	fmt.Printf("初始化已完成")
}

func p2() {
	fmt.Printf("文件对比中\n")
	file_rhash.Compare_hash()
	fmt.Printf("文件对比已完成")
}

func main() {
	// 定义命令行参数
	methodPtr := flag.String("method", "p1", "Method to call")

	// 解析命令行参数
	flag.Parse()

	// 获取命令行参数值
	method := *methodPtr

	// 检测是否存在config文件
	configFilePath := "config.json"
	exists, err := file_rhash.CheckConfigFile(configFilePath)
	if err != nil {
		fmt.Println("Error checking config file:", err)
		return
	}

	if !exists {
		fmt.Println("没有找到Config文件,正在生成中......./n")
		err = file_rhash.CreateDefaultConfig(configFilePath)
		if err != nil {
			fmt.Println("创建配置文件错误:", err)
			return
		}
		fmt.Println("已生成配置文件,请配置目录.")
		return
	}

	switch method {
	case "p1":
		p1()
	case "p2":
		p2()
	default:
		fmt.Println("Invalid method")
	}

}

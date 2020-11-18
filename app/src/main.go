package main

import (
	"fmt"
	"monitor/extend/conf"
)

func main() {
	//配置初始化
	conf.Setup()
	fmt.Println(conf.DBConf)
}

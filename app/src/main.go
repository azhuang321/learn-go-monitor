package main

import (
	"monitor/extend/conf"
	"monitor/extend/logger"
	"monitor/models"
)

func main() {
	//配置初始化
	conf.Setup()
	//日志初始化
	logger.Setup()
	//数据库的初始化
	models.Setup()

}

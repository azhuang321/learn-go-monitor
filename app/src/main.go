package main

import (
	"fmt"
	"monitor/extend/conf"
	"monitor/extend/logger"
	"monitor/extend/redis"
	"monitor/extend/validator"
	"monitor/models"
	"monitor/router"
	"monitor/schedule"
)

func main() {
	//配置初始化
	conf.Setup()
	//日志初始化
	logger.Setup()
	//数据库的初始化
	models.Setup()
	//redis初始化
	redis.Setup()
	//验证器的初始化
	validator.Setup()
	//定时任务初始化
	schedule.Setup()

	router := router.InitRouter()
	router.Run(fmt.Sprintf(":%d", conf.ServerConf.Port))
}

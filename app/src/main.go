package main

import (
	"github.com/rs/zerolog/log"
	"monitor/extend/conf"
	"monitor/extend/logger"
)

func main() {
	//配置初始化
	conf.Setup()
	//日志初始化
	logger.Setup()
	log.Info().Msg("test init")
}

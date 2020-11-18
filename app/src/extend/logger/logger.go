package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"monitor/extend/conf"
	"os"
	"strings"
)

func Setup() {
	var levelType zerolog.Level
	switch strings.ToLower(conf.LoggerConf.Level) {
	case "panic":
		levelType = zerolog.PanicLevel
	case "fatal":
		levelType = zerolog.FatalLevel
	case "error":
		levelType = zerolog.ErrorLevel
	case "warn":
		levelType = zerolog.WarnLevel
	case "info":
		levelType = zerolog.InfoLevel
	case "debug":
		levelType = zerolog.DebugLevel
	default:
		levelType = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(levelType)
	if conf.LoggerConf.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: !conf.LoggerConf.Color,
		})
	}
}

package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/ebladrocher/smarthome/system/config"
)

type ServerLogger struct {
	Logger *zap.Logger
}

func InitLogger(cfg *config.AppConfig) *ServerLogger {
	var level zapcore.Level

	var w = zapcore.AddSync(&lumberjack.Logger{
		Filename:   "var/log/logger.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	switch cfg.Mode {
		case config.ReleaseMode:
			level = zap.InfoLevel
		case config.DebugMode:
			level = zap.DebugLevel
	}

	var core = zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		level,
	)
	return &ServerLogger{Logger: zap.New(core)}
}

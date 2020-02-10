package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/ebladrocher/smarthome/system/config"
)

// Logger ...
type Logger struct {
	Logger *zap.Logger
}

// InitLogger ...
func InitLogger(cfg *config.AppConfig) *Logger {
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
	return &Logger{Logger: zap.New(core)}
}

//func (s ServerLogger) Write(b []byte) (i int, err error) {
//	s.Logger.Info(string(b))
//	return
//}

/*
cfg := zap.Config{
    Encoding:         "json",
    Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
    OutputPaths:      []string{"stderr"},
    ErrorOutputPaths: []string{"stderr"},
    EncoderConfig: zapcore.EncoderConfig{
        MessageKey: "message",

        LevelKey:    "level",
        EncodeLevel: zapcore.CapitalLevelEncoder,

        TimeKey:    "time",
        EncodeTime: zapcore.ISO8601TimeEncoder,

        CallerKey:    "caller",
        EncodeCaller: zapcore.ShortCallerEncoder,
    },
}
logger, _ = cfg.Build()
 */

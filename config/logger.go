package config

import (
	"github.com/knadh/koanf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func InitLogger(c *koanf.Koanf) {
	logConfig := zap.NewProductionConfig()
	logConfig.Sampling = nil

	var logLevel zapcore.Level

	if err := logLevel.Set(c.String("logger.level")); err != nil {
		zap.S().Fatalw("Could not determine logger level", "error", err)
	}

	logConfig.Level.SetLevel(logLevel)

	loggerEncoding := c.String("logger.encoding")
	logConfig.Encoding = loggerEncoding

	if c.Bool("logger.color") {
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Use sane timestamp when logging to console
	if logConfig.Encoding == "console" {
		logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// JSON Fields
	logConfig.EncoderConfig.MessageKey = "msg"
	logConfig.EncoderConfig.LevelKey = "level"
	logConfig.EncoderConfig.CallerKey = "caller"

	globalLogger, _ := logConfig.Build()
	zap.ReplaceGlobals(globalLogger)

}
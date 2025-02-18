package logger

import (
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"pegasus-cloud.com/aes/toolkits/mviper"
)

var _logger *zap.Logger

func InitEventLogger(fileName string) {
	eventLogger := &lumberjack.Logger{
		Filename:   filepath.Join(mviper.GetString("plugin.logger.event_log.path"), fileName),
		MaxSize:    mviper.GetInt("plugin.logger.event_log.max_size"),
		MaxBackups: mviper.GetInt("plugin.logger.event_log.max_backups"),
		MaxAge:     mviper.GetInt("plugin.logger.event_log.max_age"),
		Compress:   true,
	}
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:     "T",
		LevelKey:    "L",
		MessageKey:  "M",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime:  zapcore.ISO8601TimeEncoder,
	})
	fileWriter := zapcore.AddSync(eventLogger)
	loggerCore := zapcore.NewCore(encoder, fileWriter, zapcore.InfoLevel)
	core := zapcore.NewTee(loggerCore)
	_logger = zap.New(core, zap.AddCaller(), zap.Development())
}

func Use() *zap.Logger {
	return _logger
}

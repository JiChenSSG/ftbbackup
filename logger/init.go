package logger

import (
	"os"

	"github.com/jichenssg/ftbbackup/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() (shutdown func()) {
	writer := &lumberjack.Logger{
		Filename:   config.GetConfig().LogFile,
		MaxSize:    config.GetConfig().LogMaxSize,
		MaxBackups: config.GetConfig().LogMaxBackups,
		MaxAge:     config.GetConfig().LogMaxAge,
	}

	var level zapcore.Level
	zapconf := zap.NewProductionConfig()
	if err := level.UnmarshalText([]byte(config.GetConfig().Loglevel)); err != nil {
		level = zapcore.InfoLevel // default to info level if parsing fails
	}
	zapconf.Level = zap.NewAtomicLevelAt(level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapconf.EncoderConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(writer),
			zapcore.AddSync(os.Stdout),
		),
		zapconf.Level,
	)

	logger := zap.New(core)
	zap.ReplaceGlobals(logger)

	return func() {
		logger.Sync()
		writer.Close()
	}
}

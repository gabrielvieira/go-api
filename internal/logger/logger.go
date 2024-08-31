package logger

import (
	"github.com/gabrielvieira/go-api/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func New(config *config.Config) (*zap.Logger, error) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logLevel, err := zap.ParseAtomicLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}

	loggerConfig := zap.Config{
		Level:             logLevel,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          config.LogFormat,
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return loggerConfig.Build()
}

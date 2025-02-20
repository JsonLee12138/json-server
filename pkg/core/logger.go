package core

import (
	"github.com/JsonLee12138/json-server/pkg/configs"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"os"
	"time"
)

type Logger struct {
	config configs.LogConfig
}

// CustomTimeEncoder 自定义日志输出时间格式
func (f *Logger) CusTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(f.config.Prefix + t.Format(utils.ConvertFormat(f.config.TimeFormat)))
}

// GetEncoder 获取 zapcore.Encoder
func (f *Logger) GetEncoder() zapcore.Encoder {
	cnf := f.getEncoderConfig(f.config)
	if f.config.Format == "json" {
		return zapcore.NewJSONEncoder(cnf)
	}
	return zapcore.NewConsoleEncoder(cnf)
}

// GetEncoderCore 获取Encoder的 zapcore.Core
func (f *Logger) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := f.GetWriteSyncer(l.String()) // 日志分割
	return zapcore.NewCore(f.GetEncoder(), writer, level)
}

// GetEncoderConfig 获取zapcore.EncoderConfig
func (f *Logger) getEncoderConfig(config configs.LogConfig) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     config.MessageKey,
		LevelKey:       config.LevelKey,
		TimeKey:        config.TimeKey,
		NameKey:        config.NameKey,
		CallerKey:      config.CallerKey,
		StacktraceKey:  config.StacktraceKey,
		LineEnding:     config.LineEnding,
		EncodeLevel:    f.config.ZapEncodeLevel(),
		EncodeTime:     f.CusTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (f *Logger) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := &writeSyncer{
		config: f.config,
		level:  level,
	}
	if f.config.LogInTerminal {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (f *Logger) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := f.config.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, f.GetEncoderCore(level, GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

func NewLoggerConfig(config configs.LogConfig) *Logger {
	config.MessageKey = utils.DefaultIfEmpty(config.MessageKey, "message")
	config.LevelKey = utils.DefaultIfEmpty(config.LevelKey, "level")
	config.TimeKey = utils.DefaultIfEmpty(config.TimeKey, "time")
	config.NameKey = utils.DefaultIfEmpty(config.NameKey, "logger")
	config.CallerKey = utils.DefaultIfEmpty(config.CallerKey, "caller")
	config.LineEnding = utils.DefaultIfEmpty(config.LineEnding, zapcore.DefaultLineEnding)
	config.StacktraceKey = utils.DefaultIfEmpty(config.StacktraceKey, "stacktrace")
	config.TimeFormat = utils.DefaultIfEmpty(config.TimeFormat, "YYYY/MM/DD - HH:mm:ss")
	config.MaxBackups = utils.DefaultIfEmpty(config.MaxBackups, 10)
	config.MaxSize = utils.DefaultIfEmpty(config.MaxSize, 100)
	config.MaxAge = utils.DefaultIfEmpty(config.MaxAge, 7)
	config.Compress = utils.DefaultIfEmpty(config.Compress, true)
	config.Format = utils.DefaultIfEmpty(config.Format, "json")
	return &Logger{
		config: config,
	}
}

func NewLogger(config configs.LogConfig) (logger *zap.Logger) {
	utils.CreateDir(config.Director)
	l := NewLoggerConfig(config)
	cores := l.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	if config.ShowLineNumber {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

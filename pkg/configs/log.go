package configs

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	Director       string `mapstructure:"director" json:"director" yaml:"director" toml:"director"`
	MessageKey     string `mapstructure:"message-key" json:"message-key" yaml:"message-key" toml:"message-key"`
	LevelKey       string `mapstructure:"level-key" json:"level-key" yaml:"level-key" toml:"level-key"`
	TimeKey        string `mapstructure:"time-key" json:"time-key" yaml:"time-key" toml:"time-key"`
	NameKey        string `mapstructure:"name-key" json:"name-key" yaml:"name-key" toml:"name-key"`
	CallerKey      string `mapstructure:"caller-key" json:"caller-key" yaml:"caller-key" toml:"caller-key"`
	LineEnding     string `mapstructure:"line-ending" json:"line-ending" yaml:"line-ending" toml:"line-ending"`
	StacktraceKey  string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key" toml:"stacktrace-key"`
	Level          string `mapstructure:"level" json:"level" yaml:"level" toml:"level"`
	EncodeLevel    string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level" toml:"encode-level"`
	Prefix         string `mapstructure:"prefix" json:"prefix" yaml:"prefix" toml:"prefix"`
	TimeFormat     string `mapstructure:"time-format" json:"time-format" yaml:"time-format" toml:"time-format"`
	Format         string `mapstructure:"format" json:"format" yaml:"format" toml:"format"`
	LogInTerminal  bool   `mapstructure:"log-in-terminal" json:"logInTerminal" yaml:"log-in-terminal" toml:"log-in-terminal"`
	MaxAge         int    `mapstructure:"max-age" json:"max-age" yaml:"max-age" toml:"max-age"`
	MaxSize        int    `mapstructure:"max-size" json:"max-size" yaml:"max-size" toml:"max-size"`
	MaxBackups     int    `mapstructure:"max-backups" json:"max-backups" yaml:"max-backups" toml:"max-backups"`
	Compress       bool   `mapstructure:"" json:"cocompressmpress" yaml:"compress" toml:"compress"`
	ShowLineNumber bool   `mapstructure:"show-line-number" json:"show-line-number" yaml:"show-line-number" toml:"show-line-number"`
}

// TransportLevel 根据字符串转化为 zapcore.Level
func (c LogConfig) TransportLevel() zapcore.Level {
	c.Level = strings.ToLower(c.Level)
	switch c.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
func (c LogConfig) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case c.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case c.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case c.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case c.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

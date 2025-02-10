package global

import "github.com/JsonLee12138/json-server/pkg/configs"

type Configs struct {
	System configs.SystemConfig `mapstructure:"system" json:"system" yaml:"system" toml:"system"`
	// Mysql  configs.MysqlConfig  `mapstructure:"mysql" json:"mysql" yaml:"mysql" toml:"mysql"`
	// Redis  configs.RedisConfig  `mapstructure:"redis" json:"redis" yaml:"redis" toml:"redis"`
	//Logger core.LogConfig        `mapstructure:"logger" json:"logger" yaml:"logger"`
	//I18n   core.I18nConfig       `mapstructure:"i18n" json:"i18n" yaml:"i18n"`
	//Cors   middleware.CorsConfig `mapstructure:"cors" json:"cors" yaml:"cors"`
	//Mysql  core.MysqlConfig      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

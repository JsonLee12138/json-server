package configs

type Config struct {
	System SystemConfig `mapstructure:"system" json:"system" yaml:"system"`
	//Logger core.LogConfig        `mapstructure:"logger" json:"logger" yaml:"logger"`
	//I18n   core.I18nConfig       `mapstructure:"i18n" json:"i18n" yaml:"i18n"`
	//Cors   middleware.CorsConfig `mapstructure:"cors" json:"cors" yaml:"cors"`
	//Mysql  core.MysqlConfig      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Redis  core.RedisConfig      `mapstructure:"redis" json:"redis" yaml:"redis"`
}

package global

type SystemConfig struct {
	AppName          string `mapstructure:"app-name" json:"app-name" yaml:"app-name"`
	IPValidationAble bool   `mapstructure:"ip-validation-able" json:"ip-validation-able" yaml:"ip-validation-able"`
	RoutesPrintAble  bool   `mapstructure:"routes-print-able" json:"routes-print-able" yaml:"routes-print-able"`
	QuerySplitAble   bool   `mapstructure:"query-split-able" json:"query-split-able" yaml:"query-split-able"`
	ProxyCheckAble   bool   `mapstructure:"proxy-check-able" json:"proxy-check-able" yaml:"proxy-check-able"`
	Port             string `mapstructure:"port" json:"port" yaml:"port"`
	RouterPrefix     string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	DBType           string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
}
type Configs struct {
	System SystemConfig `mapstructure:"system" json:"system" yaml:"system"`
	//Logger core.LogConfig        `mapstructure:"logger" json:"logger" yaml:"logger"`
	//I18n   core.I18nConfig       `mapstructure:"i18n" json:"i18n" yaml:"i18n"`
	//Cors   middleware.CorsConfig `mapstructure:"cors" json:"cors" yaml:"cors"`
	//Mysql  core.MysqlConfig      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Redis  core.RedisConfig      `mapstructure:"redis" json:"redis" yaml:"redis"`
}

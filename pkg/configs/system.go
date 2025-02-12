package configs

type SystemConfig struct {
	AppName          string `mapstructure:"app-name" json:"app-name" yaml:"app-name" toml:"app-name"`
	IPValidationAble bool   `mapstructure:"ip-validation-able" json:"ip-validation-able" yaml:"ip-validation-able" toml:"ip-validation-able"`
	RoutesPrintAble  bool   `mapstructure:"routes-print-able" json:"routes-print-able" yaml:"routes-print-able" toml:"routes-print-able"`
	QuerySplitAble   bool   `mapstructure:"query-split-able" json:"query-split-able" yaml:"query-split-able" toml:"query-split-able"`
	ProxyCheckAble   bool   `mapstructure:"proxy-check-able" json:"proxy-check-able" yaml:"proxy-check-able" toml:"proxy-check-able"`
	Port             string `mapstructure:"port" json:"port" yaml:"port" toml:"port"`
	RouterPrefix     string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix" toml:"router-prefix"`
	DBType           string `mapstructure:"db-type" json:"db-type" yaml:"db-type" toml:"db-type"`
	SwaggerAble      bool   `mapstructure:"swaggerable" json:"swaggerable" yaml:"swaggerable" toml:"swaggerable"`
	ApifoxAble       bool   `mapstructure:"apifoxable" json:"apifoxable" yaml:"apifoxable" toml:"apifoxable"`
	OpenApiAble      bool   `mapstructure:"openapiable" json:"openapiable" yaml:"openapiable" toml:"openapiable"`
}

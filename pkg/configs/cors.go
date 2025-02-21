package configs

import (
	"strings"

	"github.com/JsonLee12138/jsonix/pkg/utils"
)

type CorsConfig struct {
	AllowMethods     []string `mapstructure:"allow-methods" json:"allow-methods" yaml:"allow-methods"`
	AllowHeaders     []string `mapstructure:"allow-headers" json:"allow-headers" yaml:"allow-headers"`
	AllowOrigins     []string `mapstructure:"allow-origins" json:"allow-origins" yaml:"allow-origins"`
	AllowCredentials bool     `mapstructure:"allow-credentials" json:"allow-credentials" yaml:"allow-credentials"`
	MaxAge           string   `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
}

func (c *CorsConfig) GetAllowOriginsString() string {
	return strings.Join(utils.DefaultIfEmpty(c.AllowOrigins, []string{"*"}), ",")
}

func (c *CorsConfig) GetAllowHeadersString() string {
	return strings.Join(utils.DefaultIfEmpty(c.AllowHeaders, []string{}), ",")
}

func (c *CorsConfig) GetAllowMethodsString() string {
	return strings.Join(utils.DefaultIfEmpty(c.AllowMethods, []string{}), ",")
}

func (c *CorsConfig) GetMaxAgeSeconds() int {
	str := utils.DefaultIfEmpty(c.MaxAge, "12h")
	maxAge, err := utils.ParseDuration(str)
	if err != nil {
		return 30
	}
	return int(maxAge.Seconds())
}

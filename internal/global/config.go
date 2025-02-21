package global

import "github.com/JsonLee12138/jsonix/pkg/configs"

type Configs struct {
	System  configs.SystemConfig  `mapstructure:"system" json:"system" yaml:"system" toml:"system"`
	Apifox  configs.ApifoxConfig  `mapstructure:"apifox" json:"apifox" yaml:"apifox" toml:"apifox"`
	Swagger configs.SwaggerConfig `mapstructure:"swagger" json:"swagger" yaml:"swagger" toml:"swagger"`
}

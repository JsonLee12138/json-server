package configs

import "github.com/JsonLee12138/json-server/pkg/utils"

type SwaggerConfig struct {
	BasePath string `mapstructure:"base-path" json:"base-path" yaml:"base-path" toml:"base-path"`
	FilePath string `mapstructure:"file-path" json:"file-path" yaml:"file-path" toml:"file-path"`
	Path     string `mapstructure:"path" json:"path" yaml:"path" toml:"path"`
	Title    string `mapstructure:"title" json:"title" yaml:"title" toml:"title"`
	CacheAge int    `mapstructure:"cache-age" json:"cache-age" yaml:"cache-age" toml:"cache-age"`
}

var SwaggerConfigDefault = &SwaggerConfig{
	BasePath: "/",
	FilePath: "./docs/swagger.json",
	Path:     "/swagger",
	CacheAge: 3600,
}

func (c *SwaggerConfig) Get(key string) any {
	switch key {
	case "BasePath":
		return utils.DefaultIfEmpty(c.BasePath, SwaggerConfigDefault.BasePath)
	case "FilePath":
		return utils.DefaultIfEmpty(c.FilePath, SwaggerConfigDefault.FilePath)
	case "Path":
		return utils.DefaultIfEmpty(c.Path, SwaggerConfigDefault.Path)
	case "CacheAge":
		return utils.DefaultIfEmpty(c.CacheAge, SwaggerConfigDefault.CacheAge)
	default:
		return ""
	}
}

package configs

import "fmt"

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host" toml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port" toml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password" toml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db" toml:"db"`
}

func (c *RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

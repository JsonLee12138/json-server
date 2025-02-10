package configs

import (
	"fmt"

	"github.com/JsonLee12138/json-server/pkg/utils"
)

type MysqlConfig struct {
	Loc                       string `mapstructure:"loc" json:"loc" yaml:"loc" toml:"loc"`
	Charset                   string `mapstructure:"charset" json:"charset" yaml:"charset" toml:"charset"`
	ParseTime                 string `mapstructure:"parse-time" json:"parse-time" yaml:"parse-time" toml:"parse-time"`
	Host                      string `mapstructure:"host" json:"host" yaml:"host" toml:"host"` // 数据库地址
	Port                      string `mapstructure:"port" json:"port" yaml:"port" toml:"port"`
	Config                    string `mapstructure:"config" json:"config" yaml:"config" toml:"config"`                                                                                         // 高级配置
	DBName                    string `mapstructure:"db-name" json:"db-name" yaml:"db-name" toml:"db-name"`                                                                                     // 数据库名
	Username                  string `mapstructure:"username" json:"username" yaml:"username" toml:"username"`                                                                                 // 数据库密码
	Password                  string `mapstructure:"password" json:"password" yaml:"password" toml:"password"`                                                                                 // 数据库密码
	Engine                    string `mapstructure:"engine" json:"engine" yaml:"engine" toml:"engine" default:"InnoDB"`                                                                        //数据库引擎，默认InnoDB
	LogMode                   string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode" toml:"log-mode"`                                                                                 // 是否开启Gorm全局日志
	MaxIdleConns              int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns" toml:"max-idle-conns"`                                                         // 空闲中的最大连接数
	MaxOpenConns              int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns" toml:"max-open-conns"`                                                         // 打开到数据库的最大连接数
	Singular                  bool   `mapstructure:"singular" json:"singular" yaml:"singular" toml:"singular"`                                                                                 //是否开启全局禁用复数，true表示开启
	DefaultStringSize         uint   `mapstructure:"default-string-size" json:"default-string-size" yaml:"default-string-size" toml:"default-string-size"`                                     // string 类型字段的默认长度
	SkipInitializeWithVersion bool   `mapstructure:"skip-initialize-with-version" json:"skip-initialize-with-version" yaml:"skip-initialize-with-version" toml:"skip-initialize-with-version"` // 是否根据版本自动配置，默认false
	DontSupportRenameColumn   bool   `mapstructure:"dont-support-rename-column" json:"dont-support-rename-column" yaml:"dont-support-rename-column" toml:"dont-support-rename-column"`         // 是否支持重命名列，false表示不支持, MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	DontSupportRenameIndex    bool   `mapstructure:"dont-support-rename-index" json:"dont-support-rename-index" yaml:"dont-support-rename-index" toml:"dont-support-rename-index"`             // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
}

func (c *MysqlConfig) DSN() string {
	charset := utils.DefaultIfEmpty(c.Charset, "utf8mb4")
	loc := utils.DefaultIfEmpty(c.Loc, "Local")
	parseTime := utils.DefaultIfEmpty(c.ParseTime, "True")
	if utils.IsEmpty(c.DefaultStringSize) || c.DefaultStringSize == 0 {
		c.DefaultStringSize = 256
	}
	if utils.IsEmpty(c.Engine) {
		c.Engine = "InnoDB"
	}
	if utils.IsEmpty(c.MaxIdleConns) || c.MaxIdleConns == 0 {
		c.MaxIdleConns = 50
	}
	if utils.IsEmpty(c.MaxOpenConns) || c.MaxOpenConns == 0 {
		c.MaxIdleConns = 100
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s&%s", c.Username, c.Password, c.Host, c.Port, c.DBName, charset, parseTime, loc, c.Config)
}

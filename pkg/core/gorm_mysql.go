package core

import (
	"fmt"
	"github.com/JsonLee12138/json-server/pkg/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormMysql(cnf configs.MysqlConfig, opts ...gorm.Option) *gorm.DB {
	if cnf.DBName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       cnf.DSN(),
		DefaultStringSize:         cnf.DefaultStringSize,
		DontSupportRenameIndex:    cnf.DontSupportRenameIndex,
		DontSupportRenameColumn:   cnf.DontSupportRenameColumn,
		SkipInitializeWithVersion: cnf.SkipInitializeWithVersion,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), opts...); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", fmt.Sprintf("ENGINE=%s", cnf.Engine))
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(cnf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cnf.MaxOpenConns)
		return db
	}
}

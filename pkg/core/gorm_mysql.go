package core

import (
	"errors"
	"fmt"

	"github.com/JsonLee12138/jsonix/pkg/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormMysql(cnf configs.MysqlConfig, opts ...gorm.Option) (*gorm.DB, error) {
	if cnf.DBName == "" {
		return nil, errors.New("the database name cannot be empty")
	}
	mysqlConfig := mysql.Config{
		DSN:                       cnf.DSN(),
		DefaultStringSize:         cnf.DefaultStringSize,
		DontSupportRenameIndex:    cnf.DontSupportRenameIndex,
		DontSupportRenameColumn:   cnf.DontSupportRenameColumn,
		SkipInitializeWithVersion: cnf.SkipInitializeWithVersion,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), opts...); err != nil {
		return nil, err
	} else {
		db.InstanceSet("gorm:table_options", fmt.Sprintf("ENGINE=%s", cnf.Engine))
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(cnf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cnf.MaxOpenConns)
		return db, nil
	}
}

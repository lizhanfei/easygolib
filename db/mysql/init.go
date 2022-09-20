package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql(conf *MysqlConf, logger logger.Interface) (*gorm.DB, error) {
	var err error
	var mysqlClient *gorm.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s&parseTime=True&loc=Asia%%2FShanghai",
		conf.User,
		conf.Password,
		conf.Addr,
		conf.DataBase,
		conf.ConnTimeOut,
		conf.ReadTimeOut,
		conf.WriteTimeOut)
	if "" != conf.Charset {
		dsn += "&charset=" + conf.Charset
	}
	mysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   logger,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		AllowGlobalUpdate:                        false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := mysqlClient.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)//连接池 最大空闲连接数
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)//连接池最大连接数
	sqlDB.SetConnMaxLifetime(conf.ConnMaxLifeTime)//连接池中连接最大生命周期
	sqlDB.SetConnMaxIdleTime(conf.ConnMaxIdlTime)//连接池连接最大空闲周期

	return mysqlClient, nil
}

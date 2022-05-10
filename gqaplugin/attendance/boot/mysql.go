package boot

import (
	gqaConfig "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Mysql() *gorm.DB {
	mysqlCfg := gqaConfig.Mysql{
		User:     "root",
		Password: "123456",
		Host:     "192.168.35.166",
		Port:     "3311",
		Database: "jingjigdsystem2",
		MaxIdle:  0,
		MaxOpen:  0,
	}
	dsn := mysqlCfg.User + ":" + mysqlCfg.Password + "@tcp(" + mysqlCfg.Host + ":" + mysqlCfg.Port + ")/" + mysqlCfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(mysqlCfg.MaxIdle)
		sqlDB.SetMaxOpenConns(mysqlCfg.MaxOpen)
		return db
	}
}

func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, NamingStrategy: schema.NamingStrategy{SingularTable: true}}
	return config
}

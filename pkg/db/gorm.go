package db

import (
	"mall-pkg/config"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//@author: SliverHorn
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config
func gormConfig(c *config.Mysql, log *zap.Logger) *gorm.Config {
	cfg := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	}
	switch c.LogMode {
	case "silent", "Silent":
		cfg.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		cfg.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		cfg.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		cfg.Logger = Default.LogMode(logger.Info)
	default:
		cfg.Logger = Default.LogMode(logger.Info)
	}
	return cfg
}

// 新建GORM数据库实例
func NewGorm(c config.Mysql, log *zap.Logger) (*gorm.DB, error) {
	dsn := c.Username + ":" + c.Password + "@tcp(" + c.Path + ")/" + c.Dbname + "?" + c.Config
	cfg := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	mdb, err := gorm.Open(mysql.New(cfg), gormConfig(&c, log))
	if err != nil {
		return nil, err
	}

	sqlDB, _ := mdb.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	// db.Use(dbresolver.Register(dbresolver.Config{
	// 	Replicas: []gorm.Dialector{mysql.Open("root:m156018!@#@tcp(127.0.0.1)/game?charset=utf8mb4&parseTime=True&loc=Local")},
	// }, "user", "wallet"))
	return mdb, nil

}

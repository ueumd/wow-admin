package utils

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"wow-admin/config"
	"wow-admin/model"
)

func InitMyQLDB() *gorm.DB {
	mysqlCfg := config.Cfg.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.Username,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), gormConfig())

	if err != nil {
		log.Fatal("MySQL 连接失败, 请检查参数")
	}

	log.Println("MySQL 连接成功")

	// 注册数据库表专用
	err = db.AutoMigrate(
		&model.City{},

		//RBAC
		&model.Role{},
		&model.Resource{},
		&model.Menu{},
		&model.UserAuth{},
		&model.RoleResource{},
		&model.RoleMenu{},
		&model.UserRole{},
	)
	if err != nil {
		Logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	sqlDB, _ := db.DB()

	// 设置连接池中的最大闲置连接
	sqlDB.SetMaxIdleConns(10)

	// 设置数据库的最大连接数量
	sqlDB.SetMaxOpenConns(100)

	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return db
}

func gormConfig() *gorm.Config {
	return &gorm.Config{
		// gorm 日志模式
		Logger: logger.Default.LogMode(getLogMode(config.Cfg.Mysql.LogMode)),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	}
}

// 根据字符串获取对应 LogLevel
func getLogMode(str string) logger.LogLevel {
	switch str {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}

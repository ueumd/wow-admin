package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 注册MySQL驱动
	"github.com/jmoiron/sqlx"
	"time"
	"wow-admin/config"
	"wow-admin/global"
)

func InitSqlxDB() (err error) {
	mysqlCfg := config.Cfg.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.Username,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Dbname,
	)

	//建立链接db, err := sql.Open("mysql", "user:password@/dbname")
	global.DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("Open Connection failed:", err.Error()))
	}

	// 设置连接池中的最大闲置连接
	global.DB.SetMaxIdleConns(10)

	// 设置数据库的最大连接数量
	global.DB.SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	global.DB.SetConnMaxLifetime(10 * time.Second)

	return
}

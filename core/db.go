package core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 注册MySQL驱动
	"github.com/jmoiron/sqlx"
	"wow-admin/global"
)

func InitDB() error {
	var err error
	cfg := global.CONFIG.DB
	//建立链接db, err := sql.Open("mysql", "user:password@/dbname")
	global.DB, err = sqlx.Open("mysql",cfg.Address)
	if err != nil {
		panic(	fmt.Errorf("Open Connection failed:", err.Error()))
	}
	global.DB.SetMaxIdleConns(cfg.MaxIdles)
	global.DB.SetMaxOpenConns(cfg.MaxConns)
	return global.DB.Ping()
}
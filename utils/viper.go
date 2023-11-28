package utils

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"wow-admin/config"
)

func InitViper() {
	var configPath string
	flag.StringVar(&configPath, "c", "", "choose config file")
	flag.Parse()

	if configPath != "" {
		log.Printf("命令行读取参数, 配置文件路径为: %s\n", configPath)
	} else {
		configPath = "config/config.yml"
	}

	v := viper.New()

	// 设置路径
	v.SetConfigFile(configPath)

	// 允许使用环境变量
	v.AutomaticEnv()

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败: ", err)
	}

	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件内容加载失败: ", err)
	}

	log.Println("配置文件加载成功")
	// log.Println(config.Cfg)
}

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var _cfg *Config

type Config struct {
	Port	int 			`yaml:"port"`
	DB		dbConfig		`yaml:"dbConfig"`
	Redis	redisConfig		`yaml:"redisConfig"`
	Zap		zapConfig		`yaml:"zap"`
	Log		logConfig		`yaml:"logConfig"`
}

type dbConfig struct {
	Address		string		`json:"address" yaml:"address"`
	MaxConns 	int    		`json:"maxConn" yaml:"maxConn"` // 空闲中的最大连接数
	MaxIdles 	int    		`json:"minIdle" yaml:"minIdle"` // 打开到数据库的最大连接数
}

type redisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	PoolSize int    `yaml:"poolSize"`
}

type zapConfig struct {
	Level        string `json:"level" yaml:"level"`
	Filename     string `json:"filename" yaml:"filename"`
	MaxSize      int    `json:"maxSize" yaml:"maxSize"`
	MaxBackups   int    `json:"maxBackups" yaml:"maxBackups"`
	MaxAge       int    `json:"maxAge" yaml:"maxAge"`
	LogInConsole bool   `json:"logInConsole" yaml:"log-in-console"` // 输出控制台
}

type logConfig struct {
	FilePath string `yaml:"filePath"`
	IsStdOut bool   `yaml:"isStdout"`
	LogLevel uint32 `yaml:"logLevel"`
}

func Init(filePath string) error {
	bys, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	_cfg = &Config{}
	err = yaml.Unmarshal(bys, _cfg)
	return err
}

func Get() Config {
	if _cfg == nil {
		log.Fatal("config file is not init")
	}

	return *_cfg
}

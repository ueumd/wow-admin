package config

/*
*
配置文件
*/
var Cfg Config

type Config struct {
	Server  Server
	JWT     JWT
	Mysql   Mysql
	Redis   Redis
	Session Session
	Zap     Zap
}

type Server struct {
	AppMode   string
	BackPort  string
	FrontPort string
}

type JWT struct {
	Secret string // JWT 签名
	Expire int64  // 过期时间
	Issuer string // 签发者
}

type Mysql struct {
	Host     string // 服务器地址
	Port     string // 端口
	Config   string // 高级配置
	Dbname   string // 数据库名
	Username string // 数据库用户名
	Password string // 数据库密码
	LogMode  string // 日志级别
}

type Redis struct {
	Address  string
	Password string
	DB       int
}

type Zap struct {
	Level        string // 级别
	Prefix       string // 日志前缀
	Format       string // 输出
	Directory    string // 日志文件夹
	MaxAge       int    // 日志留存时间
	ShowLine     bool   // 显示行
	LogInConsole bool   // 输出控制台
}

type Session struct {
	Name   string
	Salt   string
	MaxAge int
}

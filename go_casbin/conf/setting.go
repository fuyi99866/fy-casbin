package conf

import (
	"github.com/go-ini/ini"
	"go_casbin/pkg/logger"
	"time"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type SWAGGER struct {
	Host string
}

var Swag = &SWAGGER{}

type SERVER struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	HTTPS        bool
	BasePath     string
}

var ServerSetting = &SERVER{}

type App struct {
	JwtSecret   string
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	LogLever    string
	TimeFormat  string
}

var AppSetting = &App{}

var cfg *ini.File

func Init(config string) {
	var err error
	cfg, err = ini.Load(config)
	if err != nil {
		logger.Fatal("初始化配置文件失败： ", err)
	}
	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)
	mapTo("swagger", Swag)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logger.Fatalf("加载配置文件失败", err)
	}
}

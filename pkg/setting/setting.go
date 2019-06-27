package setting

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

type App1 struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = NewDefaultAppSetting()

var ServerSetting = &Server{}

var DatabaseSetting = &Database{}

var RedisSetting = &Redis{}

var CacheSetting = &Cache{}

var cfg *ini.File

func Setup() {
	var err error
	dir, _ := os.Getwd()
	path := dir + "/conf/app.ini"
	cfg, err = ini.Load(path)
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("cache", CacheSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

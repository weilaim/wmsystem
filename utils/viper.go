package utils

import (
	"flag"
	"log"
	"strings"

	"github.com/spf13/viper"
	"github.com/weilaim/wmsystem/config"
)

func InitViper() {
	// 根据命令行读取配置文件路径
	var configPath string
	flag.StringVar(&configPath, "c", "", "chosse config file.")
	flag.Parse()
	if configPath != "" { // 命令行读取到参数
		log.Printf("命令行读取参数，配置文件路径为:%s\n", configPath)
	} else { // 命令行未读取到参数
		log.Println("命令行参数为空，默认加载:config/config.toml")
		configPath = "config/config.toml"
	}

	// 读取固定路径文件配置 config/config.toml
	v := viper.New()
	v.SetConfigFile(configPath)
	v.AutomaticEnv()                                   // 允许使用环境变量
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // SERVER_APPMODE => SERVER.APPMODE

	// 读取配置InitViper
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败:", err)
	}

	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件内容加载失败:", err)
	}

	log.Println("配置文件内容加载成功")
}

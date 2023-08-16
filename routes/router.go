package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/weilaim/wmsystem/config"
	"github.com/weilaim/wmsystem/dao"
	"github.com/weilaim/wmsystem/utils"
)

// 初始化全局变量
func InitGlobalVariable() {
	// 初始化viper
	utils.InitViper()
	// 初始化Logger
	utils.InitLogger()

	// 初始化数据库 DB
	dao.DB = utils.InitMySQLDB()

	// 初始化Redis
	utils.InitRedis()

	// 初始化Casbin
	utils.InitCasbin(dao.DB)

}

// 后台接口服务
func BackendServer() *http.Server {
	backPort := config.Cfg.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      BackRouter(), // 后台接口管理 
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

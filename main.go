package main

import (
	"log"

	"github.com/weilaim/wmsystem/dao"
	"github.com/weilaim/wmsystem/routes"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	// 初始化全局变量
	routes.InitGlobalVariable()

	// 后台接口服务
	g.Go(func() error {
		return routes.BackendServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	if dao.DB != nil {
		// 程序结束前关闭数据库连接
		db, _ := dao.DB.DB()
		defer db.Close()
	}
}

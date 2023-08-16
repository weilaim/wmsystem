package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/weilaim/wmsystem/config"
	"github.com/weilaim/wmsystem/routes/middleware"
)

// 后台管理页面的接口路由
func BackRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	r.SetTrustedProxies([]string{"*"})

	r.Use(middleware.Logger())             // 自定义的zap日志中间件
	r.Use(middleware.ErrorRecovery(false)) // 自定义错误处理中间件
	r.Use(middleware.Cors())               // 跨域中间件

	// 基于 cookie 存储session
	store := cookie.NewStore([]byte(config.Cfg.Session.Salt))

	// session 存储时间跟jwt过期时间一致
	store.Options(sessions.Options{MaxAge: int(config.Cfg.JWT.Expire) * 3600})
	r.Use(sessions.Sessions(config.Cfg.Session.Name, store)) // Session 中间件

	// 无需鉴权的接口
	base := r.Group("/api")
	{
		// TODO 用户注册和后台登录应该记录导日志
		base.POST("/login", userAuthAPI.Login)
	}
	return r
}

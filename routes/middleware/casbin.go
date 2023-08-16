package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/wmsystem/utils"
	"github.com/weilaim/wmsystem/utils/r"
)

// casbin 鉴权中间件
func RBAC() gin.HandlerFunc {
	// 重新加载策略 确保是最新的策略
	utils.Casbin.LoadPolicy()
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		url, method := c.FullPath()[4:], c.Request.Method
		// 使用casbin 自带的验证函数执行策略
		fmt.Println("=======> casbin 权限管理：", role, url, method)
		isPass, err := utils.Casbin.Enforcer().Enforce(role, url, method)
		// 权限严重未通过
		if err != nil || !isPass {
			r.SendCode(c, r.ERROR_PERMI_DENIED)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

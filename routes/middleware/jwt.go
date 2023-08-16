package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/wmsystem/utils"
	"github.com/weilaim/wmsystem/utils/r"
)

// JWT
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 约定token放在 Header 的 Authorization 中， 并使用Bearer开头
		token := c.Request.Header.Get("Authorization")
		// token nil
		if token == "" {
			r.SendCode(c, r.ERROR_TOKEN_NOT_EXIST)
			c.Abort()
			return
		}

		// token 的正确格式
		parts := strings.Split(token, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			r.SendCode(c, r.ERROR_TOKEN_TYPE_WRONG)
			c.Abort()
			return
		}

		// parts[1] 是获取到的 tokenString  使用JWT 解析函数解析它
		caims, err := utils.GetJWT().ParseToken(parts[1])
		if err != nil {
			r.SendData(c, r.ERROR_TOKEN_WRONG, err.Error())
			c.Abort()
			return
		}

		// 判断token过期时间
		if time.Now().Unix() > caims.ExpiresAt.Unix() {
			r.SendCode(c, r.ERROR_TOKEN_RUNTIME)
			c.Abort()
			return
		}

		// 将当前请求的相关信息保存到请求的上下文 c 上
		// 后续处理函数可用 c.Get("xxx") 来获取当前的请求的用户信息
		c.Set("user_info_id", caims.UserId)
		c.Set("role", caims.Role)
		c.Set("uuid", caims.UUID)
		c.Next()
	}
}

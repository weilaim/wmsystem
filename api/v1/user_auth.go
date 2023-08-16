package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/wmsystem/model/req"
	"github.com/weilaim/wmsystem/utils"
	"github.com/weilaim/wmsystem/utils/r"
)

type UserAuth struct{}

func (*UserAuth) Register(c *gin.Context) {
	r.SendCode(c, userService.Register(utils.BindValidJson[req.Register](c)))
}

func (*UserAuth) Login(c *gin.Context) {
	loginReq := utils.BindValidJson[req.Login](c)
	loginVo, code := userService.Login(c, loginReq.Username, loginReq.Password)
	r.SendData(c, code, loginVo)
}

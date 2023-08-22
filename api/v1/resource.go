package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/wmsystem/model/req"
	"github.com/weilaim/wmsystem/utils"
	"github.com/weilaim/wmsystem/utils/r"
)

type Resource struct{}

func (*Resource) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, resourceService.SaveOrUpdate(utils.BindJson[req.SaveOrUpdateResource](c)))
}

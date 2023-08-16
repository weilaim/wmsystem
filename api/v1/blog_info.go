package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BlogInfo struct{}

func (*BlogInfo) GetHomeInfo(c *gin.Context) {
	fmt.Println("houtaishouyeinxi")
}

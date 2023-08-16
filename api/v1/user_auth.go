package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserAuth struct{}

func (*UserAuth) Login(c *gin.Context) {
	fmt.Println("登录啦")
}

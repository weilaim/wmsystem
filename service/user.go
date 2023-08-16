package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/wmsystem/dao"
	"github.com/weilaim/wmsystem/model"
	"github.com/weilaim/wmsystem/model/req"
	"github.com/weilaim/wmsystem/model/resp"
	"github.com/weilaim/wmsystem/utils"
	"github.com/weilaim/wmsystem/utils/r"
)

type User struct{}

// login
func (*User) Login(c *gin.Context, username, password string) (loginVo resp.LoginVo, code int) {
	fmt.Println("service login")
	return loginVo, 1
}

// Register

func (*User) Register(req req.Register) (code int) {
	// 检查验证码是否正确
	// if req.Code != utils.Redis.GetVal(KEY_CODE+req.Username) {
	// 	return r.ERROR_VERIFICATION_CODE
	// }

	// 检查用户名是否已存在，则该账号已经注册过
	if exist := checkUserExistByName(req.Username); exist {
		return r.ERROR_USER_NAME_USED
	}

	userInfo := &model.UserInfo{
		Email:    req.Username,
		Nickname: "用户" + req.Username,
		Avatar:   blogInfoService.GetBlogConfig().UserAvatar,
	}

	dao.Create(&userInfo)
	// 设置默认角色
	dao.Create(&model.UserRole{
		UserId: userInfo.ID,
		RoleId: 3, //默认角色是" 测试 "
	})

	dao.Create(&model.UserAuth{
		UserInfoId:    userInfo.ID,
		Username:      req.Username,
		Password:      utils.Encryptor.BcryptHash(req.Password),
		LoginType:     1,
		LastLoginTime: time.Now(), // 注册时间会更新 “上次登录时间”
	})

	return r.OK
}

func checkUserExistByName(username string) bool {
	exisUser := dao.GetOne(model.UserAuth{}, "username = ?", username)
	fmt.Println(exisUser)
	return exisUser.ID != 0
}

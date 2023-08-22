package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/weilaim/wmsystem/config"
	"github.com/weilaim/wmsystem/dao"
	"github.com/weilaim/wmsystem/model"
	"github.com/weilaim/wmsystem/model/dto"
	"github.com/weilaim/wmsystem/model/req"
	"github.com/weilaim/wmsystem/model/resp"
	"github.com/weilaim/wmsystem/utils"
	"github.com/weilaim/wmsystem/utils/r"
	"go.uber.org/zap"
)

type User struct{}

// login
func (*User) Login(c *gin.Context, username, password string) (loginVo resp.LoginVo, code int) {
	// 检查用户是否存在
	userAuth := dao.GetOne(model.UserAuth{}, "username", username)
	if userAuth.ID == 0 {
		return loginVo, r.ERROR_USER_NOT_EXIST
	}

	// 检查密码是否正确
	if !utils.Encryptor.BcryptCheck(password, userAuth.Password) {
		return loginVo, r.ERROR_PASSWORD_WRONG
	}

	// 获取用户详情信息 DTO
	userDetailDTO := convertUserDetailDTO(userAuth, c)

	// 登录信息正确， 生成token
	// TODO: 目前只给用户设定一个角色，获取第一个值就行，后期优化：给用户设置多个角色
	// uuid 生产方法 ip + 浏览器信息 + 操作系统信息
	uuid := utils.Encryptor.MD5(userDetailDTO.IpAddress + userDetailDTO.Browser + userDetailDTO.OS)
	token, err := utils.GetJWT().GenToken(userAuth.ID, userDetailDTO.RoleLabels[0], uuid)
	if err != nil {
		utils.Logger.Info("登录时生成Token错误：", zap.Error(err))
		return loginVo, r.ERROR_TOKEN_CREATE
	}

	userDetailDTO.Token = token
	// 更新用户验证信息：ip 信息 + 上次登录信息
	dao.Update(&model.UserAuth{
		Universal:     model.Universal{ID: userAuth.ID},
		IpAddress:     userDetailDTO.IpAddress,
		IpSource:      userDetailDTO.IpSource,
		LastLoginTime: userAuth.LastLoginTime,
	}, "ip_address", "ip_source", "last_login_time")

	// 保存用户到session中和redis中
	session := sessions.Default(c)
	// session只能存字符串
	sessionInfoStr := utils.Json.Marshal(dto.SessionInfo{UserDetailDTO: userDetailDTO})
	session.Set(KEY_USER+uuid, sessionInfoStr)
	utils.Redis.Set(KEY_USER+uuid, sessionInfoStr, time.Duration(config.Cfg.Session.MaxAge)*time.Second)
	session.Save()

	return userDetailDTO.LoginVo, r.OK
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

// 转化UserDetailDTO
func convertUserDetailDTO(userAuth model.UserAuth, c *gin.Context) dto.UserDetailDTO {
	ipAddress := utils.IP.GetIpAddress(c)
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress)
	// ipSource := "东莞" // fmt.Println(ipSource, "ipsource")
	browser, os := "unknown", "unknown"

	if userAgent := utils.IP.GetUserAgent(c); userAgent != nil {
		browser = userAgent.Name + " " + userAgent.Version.String()
		os = userAgent.OS + " " + userAgent.OSVersion.String()
	}

	// 获取用户详情信息
	userInfo := dao.GetOne(&model.UserInfo{}, "id", userAuth.ID)
	// 获取用户角色，没有角色默认是"test"
	roleLabels := roleDao.GetLabelsByUserInfoId(userInfo.ID)
	if len(roleLabels) == 0 {
		roleLabels = append(roleLabels, "test")
	}

	// 用户点赞 set
	articleLikeSet := utils.Redis.SMembers(KEY_ARTICLE_USER_LIKE_SET + strconv.Itoa(userInfo.ID))
	commentLikeSet := utils.Redis.SMembers(KEY_COMMENT_USER_LIKE_SET + strconv.Itoa(userInfo.ID))

	return dto.UserDetailDTO{
		LoginVo: resp.LoginVo{
			ID:             userAuth.ID,
			UserInfoId:     userInfo.ID,
			Email:          userInfo.Email,
			LoginType:      userAuth.LoginType,
			Username:       userAuth.Username,
			Nickname:       userInfo.Nickname,
			Avatar:         userInfo.Avatar,
			Intro:          userInfo.Intro,
			Website:        userInfo.Website,
			IpAddress:      ipAddress,
			IpSource:       ipSource,
			LastLoginTime:  time.Now(),
			ArticleLikeSet: articleLikeSet,
			CommentLikeSet: commentLikeSet,
		},
		Password:   userAuth.Password,
		RoleLabels: roleLabels,
		IsDisable:  userInfo.IsDisable,
		Browser:    browser,
		OS:         os,
	}
}

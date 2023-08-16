package service

import (
	"github.com/weilaim/wmsystem/dao"
	"github.com/weilaim/wmsystem/model"
	"github.com/weilaim/wmsystem/utils"
)

type BlogInfo struct{}

func (*BlogInfo) GetBlogConfig() (respVo model.BlogConfigDetail) {
	// 尝试从Redis 中取值
	blogConfig := utils.Redis.GetVal(KEY_BLOG_CONFIG)
	// Redis 中没有值，在查数据库，查到后设置到redis中
	if blogConfig == "" {
		blogConfig = dao.GetOne(model.BlogConfig{}, "id", 1).Config
		utils.Redis.Set(KEY_BLOG_CONFIG, blogConfig, 0)
	}

	utils.Json.Unmarshal(blogConfig, &respVo)
	return respVo
}

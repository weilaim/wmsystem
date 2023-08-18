package service

import "github.com/weilaim/wmsystem/dao"

const (
	KEY_CODE        = "code"        // 验证码
	KEY_USER        = "user"        // 记录用户
	KEY_BLOG_CONFIG = "blog_config" // 博客配置信息

	KEY_ARTICLE_USER_LIKE_SET = "article_user_like:" // 文章点赞Set

	KEY_COMMENT_USER_LIKE_SET = "comment_user_like:" // 评论点赞 Set
)

var (
	roleDao dao.Role
)

var (
	blogInfoService BlogInfo
)

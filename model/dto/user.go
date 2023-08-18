package dto

import "github.com/weilaim/wmsystem/model/resp"

type UserDetailDTO struct {
	resp.LoginVo
	Password   string `json:"password"`
	IsDisable  int8   `json:"is_disable"`
	Browser    string `json:"browser"`
	OS         string `json:"os"`
	RoleLabels []string `json:"role_labels"`
}


// session 信息：记录用户详细信息+是否被强退 
type SessionInfo struct {
	UserDetailDTO
	IsOffline int `json:"is_offline"`
}

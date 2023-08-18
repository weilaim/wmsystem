package dao

type Role struct{}

// 根据[userInfoId] 获取角色标签列表
func (*Role) GetLabelsByUserInfoId(userInfoId int) (labels []string) {
	DB.Table("role r, user_role ur").
		Where("r.id = ur.role_id AND ur.user_id = ?", userInfoId).
		Pluck("label", &labels) // 将单列查询为切片

	return
}


// 

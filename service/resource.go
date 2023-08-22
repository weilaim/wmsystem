package service

import (
	"fmt"

	"github.com/weilaim/wmsystem/dao"
	"github.com/weilaim/wmsystem/model"
	"github.com/weilaim/wmsystem/model/req"
	"github.com/weilaim/wmsystem/utils"
	"github.com/weilaim/wmsystem/utils/r"
)

type Resource struct{}

func (*Resource) SaveOrUpdate(req req.SaveOrUpdateResource) (code int) {
	// 检查资源名是否已存在
	existByName := dao.GetOne(model.Resource{}, "name", req.Name)

	if existByName.ID != 0 && existByName.ID != req.ID {
		return r.ERROR_RESOURCE_NAME_EXIST
	}

	fmt.Println(req.ID)
	// 更新 or 新增
	if req.ID != 0 { // 更新
		fmt.Println("req.id !=0")

	} else { // 新增
		data := utils.CopyProperties[model.Resource](req)
		dao.Create(&data)
	}

	return 0
}

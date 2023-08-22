package utils

import (
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

// 拷贝属性  vo->po
func CopyProperties[T any](from any) (to T) {
	if err := copier.Copy(&to, from); err != nil {
		Logger.Error("CopyProperties: ", zap.Error(err))
		panic(err)
	}

	return to
}

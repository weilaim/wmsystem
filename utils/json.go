package utils

import (
	"encoding/json"

	"go.uber.org/zap"
)

const JSON_UTIL_ERR_PREFIX = "utils/json.go ->"

var Json = new(_json)

type _json struct {
}

// jsonStr -> data
func (*_json) Unmarshal(data string, v any) {
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		Logger.Error(JSON_UTIL_ERR_PREFIX+"Unmarshal: ", zap.Error(err))
		panic(err)
	}
}

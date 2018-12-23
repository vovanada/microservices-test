package helpers

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"strconv"
)

func GetIntParam(req *restful.Request, param string, defaultValue int64) (int64, error) {
	value := req.QueryParameter(param)

	if value == "" {
		return defaultValue, nil
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return intValue, fmt.Errorf("param [%s] is not int, val: %s", param, err)
	}

	return intValue, nil
}

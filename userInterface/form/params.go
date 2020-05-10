package form

import (
	"github.com/labstack/echo"
	"golang.org/x/xerrors"
	"strconv"
	"encoding/json"
)

type Params struct {
	Name    string
	Page    int64
	Perpage int64
}


func GetParams(c echo.Context) (params Params, err error) {
	params, err = ParameterMapping(c)
	if err != nil {
		return params, xerrors.Errorf("can not get parameters: %w", err)
	}
	return params, nil
}

func ParameterMapping(c echo.Context) (params Params, err error) {
	parammap := c.QueryParams()
	intkeyarray := []string{"page", "perpage", "enterprise_id"}
	data := make(map[string]interface{})
	for k, v := range parammap {
		if Contain(intkeyarray, k) {
			data[k], _ = strconv.ParseInt(v[0], 10, 64)
		} else {
			data[k] = string(v[0])
		}
	}
	paramjson, err := json.Marshal(data)
	if err != nil {
		return params, err
	}
	json.Unmarshal([]byte(paramjson), &params)
	return params, nil
}

func Contain(arr []string, str string) bool{
  for _, v := range arr{
    if v == str{
      return true
    }
  }
  return false
}
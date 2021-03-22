package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

//从gin框架的body中读取数据，防止读取一次后报错
func GetFromGinBody(c *gin.Context, data interface{}) error{
	//非get请求从body中获取c_id
	body, err := ioutil.ReadAll(c.Request.Body)
	//将数据重新写入body，因为readall读取一次后就不在了，会导致后面的api服务读取body是报错"unexpected EOF"
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if err != nil {
		return errors.New("读取body失败")
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return err
	}
	return nil
}

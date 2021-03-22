package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type BaseController struct {
	ReqParams []byte
}

//准备，读取请求变量
func (bc *BaseController) Prepare(c *gin.Context) {
	bc.ReqParams, _ = ioutil.ReadAll(c.Request.Body)
}

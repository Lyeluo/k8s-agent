package filter

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"k8s.agent/pkg/model"
)

const (
	SECRET_KEY = "secretKey"
)

// c.Next():中断当前函数 但是Use调用链中其他的函数以及rest请求会继续进行
// c.Abort()： 中断当前函数，并且调用链在执行完当前函数后不会继续执行其他函数
func MiddlewareAuth(c *gin.Context) {

	if !viper.GetBool("auth.open") {
		c.Next()
		return
	}
	secretKey := c.GetHeader(SECRET_KEY)

	if secretKey == "" || secretKey != getSecretValue() {
		// 日志使用方法
		zap.L().Error("调用接口校验未通过")
		c.Abort()
		c.JSON(http.StatusUnauthorized, model.NewResponse(false, nil, model.AuthErr.Msg))
		return
	}
	c.Next()
}

// md5加密
func getSecretValue() (md5str string) {
	secretToken := viper.GetString("auth.secretToken")
	data := []byte(secretToken)
	has := md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	return
}

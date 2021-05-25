package filter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/model"
)

const (
	SECRET_KEY = "secretKey"
)

func MiddlewareAuth(c *gin.Context) {
	secretKey := c.GetHeader(SECRET_KEY)

	if secretKey == "" {
		// 日志使用方法
		zap.L().Error("调用接口校验未通过")
		c.JSON(http.StatusUnauthorized, model.NewResponse(false, nil, model.AuthErr.Msg))

		return
	}
	c.Next()
}

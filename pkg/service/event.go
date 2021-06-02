package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	myutil "k8s.agent/pkg/util"
)

// 查询当前命名空间下的所有service
func EventList(c *gin.Context) {
	namespace := c.Param("namespace")

	listOptions, err := myutil.GetListOptions(c)
	if err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	eventList, err := config.GetK8sClient().CoreV1().Events(namespace).List(listOptions)
	if err != nil {
		zap.L().Sugar().Errorf("查询event失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, eventList, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, eventList, model.NoErr.Msg))

}

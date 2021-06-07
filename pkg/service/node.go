package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	myutil "k8s.agent/pkg/util"
	"k8s.io/apimachinery/pkg/types"
)

// 查询所有的node节点
func NodeList(c *gin.Context) {

	listOptions, err := myutil.GetListOptions(c)
	if err != nil {
		zap.L().Sugar().Errorf("查询node节点信息失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	nodeList, err := config.GetK8sClient().CoreV1().Nodes().List(listOptions)
	if err != nil {
		zap.L().Sugar().Errorf("查询node节点信息失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nodeList, model.NoErr.Msg))

}

// 更改node节点信息（只允许修改部分注释和标签）
func NodeUpdate(c *gin.Context) {
	name := c.Param("name")

	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		zap.L().Sugar().Errorf("更新node失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	playLoadBytes, _ := json.Marshal(data)
	nodeResult, err := config.GetK8sClient().CoreV1().Nodes().Patch(name, types.StrategicMergePatchType, playLoadBytes)
	if err != nil {
		zap.L().Sugar().Errorf("更新node失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nodeResult, model.NoErr.Msg))

}

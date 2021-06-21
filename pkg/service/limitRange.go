package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 查询LimitRange列表
func LimitRangeList(c *gin.Context) {
	namespace := c.Param("namespace")

	limitRangeList, err := config.GetK8sClient().CoreV1().LimitRanges(namespace).List(metav1.ListOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询LimitRange失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, limitRangeList, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, limitRangeList, model.NoErr.Msg))
}

// 根据名称查询LimitRange
func LimitRangeGet(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	limitRange, err := config.GetK8sClient().CoreV1().LimitRanges(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询LimitRange失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, limitRange, model.NoErr.Msg))
}

// 创建LimitRange
func LimitRangeCreate(c *gin.Context) {
	LimitRange := v1.LimitRange{}
	namespace := c.Param("namespace")
	err := c.BindJSON(&LimitRange)
	if err != nil {
		zap.L().Sugar().Errorf("创建LimitRange失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	limitRange, err := config.GetK8sClient().CoreV1().LimitRanges(namespace).Create(&LimitRange)
	if err != nil {
		zap.L().Sugar().Errorf("创建LimitRange失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, limitRange, model.NoErr.Msg))
}

// 修改LimitRange
func LimitRangeUpdate(c *gin.Context) {
	namespace := c.Param("namespace")
	LimitRange := v1.LimitRange{}

	err := c.BindJSON(&LimitRange)
	if err != nil {
		zap.L().Sugar().Errorf("修改LimitRange失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	limitRange, err := config.GetK8sClient().CoreV1().LimitRanges(namespace).Update(&LimitRange)
	if err != nil {
		zap.L().Sugar().Errorf("修改LimitRange失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, limitRange, model.NoErr.Msg))

}

// 删除LimitRange
func LimitRangeDelete(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	err := config.GetK8sClient().CoreV1().LimitRanges(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除LimitRange失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nil, model.NoErr.Msg))
}

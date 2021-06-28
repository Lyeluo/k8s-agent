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

// 查询Quota列表
func QuotaList(c *gin.Context) {
	namespace := c.Param("namespace")

	QuotaList, err := config.GetK8sClient().CoreV1().ResourceQuotas(namespace).List(metav1.ListOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询Quota失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, QuotaList, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, QuotaList, model.NoErr.Msg))
}

// 根据名称查询Quota
func QuotaGet(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	Quota, err := config.GetK8sClient().CoreV1().ResourceQuotas(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询Quota失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, Quota, model.NoErr.Msg))
}

// 创建Quota
func QuotaCreate(c *gin.Context) {
	Quota := v1.ResourceQuota{}
	namespace := c.Param("namespace")
	err := c.BindJSON(&Quota)
	if err != nil {
		zap.L().Sugar().Errorf("创建Quota失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	QuotaRes, err := config.GetK8sClient().CoreV1().ResourceQuotas(namespace).Create(&Quota)
	if err != nil {
		zap.L().Sugar().Errorf("创建Quota失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, QuotaRes, model.NoErr.Msg))
}

// 修改Quota
func QuotaUpdate(c *gin.Context) {
	namespace := c.Param("namespace")
	Quota := v1.ResourceQuota{}

	err := c.BindJSON(&Quota)
	if err != nil {
		zap.L().Sugar().Errorf("修改Quota失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	QuotaRes, err := config.GetK8sClient().CoreV1().ResourceQuotas(namespace).Update(&Quota)
	if err != nil {
		zap.L().Sugar().Errorf("修改Quota失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, QuotaRes, model.NoErr.Msg))
}

// 删除Quota
func QuotaDelete(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	err := config.GetK8sClient().CoreV1().ResourceQuotas(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除Quota失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nil, model.NoErr.Msg))
}

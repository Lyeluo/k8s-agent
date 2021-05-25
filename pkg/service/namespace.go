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

func NamespaceList(c *gin.Context) {

	ns, err := config.GetK8sConfig().CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询namespace失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, ns, model.NoErr.Msg))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, ns, model.NoErr.Msg))
}

func NamespaceGet(c *gin.Context) {
	name := c.Param("name")
	ns, err := config.GetK8sConfig().CoreV1().Namespaces().Get(name, metav1.GetOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询namespace失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, ns, model.NoErr.Msg))
}

func NamespaceCreate(c *gin.Context) {
	var namespace *v1.Namespace

	err := c.BindJSON(namespace)
	if err != nil {
		zap.L().Sugar().Errorf("创建namespace失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, model.NoErr.Msg))
		return
	}

	ns, err := config.GetK8sConfig().CoreV1().Namespaces().Create(namespace)
	if err != nil {
		zap.L().Sugar().Errorf("创建namespace失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, model.NoErr.Msg))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, ns, model.NoErr.Msg))
}

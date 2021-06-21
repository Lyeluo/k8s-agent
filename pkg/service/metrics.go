package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	myutil "k8s.agent/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// 根据node名称查询资源使用信息
func MetricsNode(c *gin.Context) {

	mc, err := metrics.NewForConfig(config.GetK8sConfig())
	if err != nil {
		panic(err)
	}
	listOptions, err := myutil.GetListOptions(c)
	if err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	nodeMetrics, err := mc.MetricsV1beta1().NodeMetricses().List(listOptions)
	if err != nil {
		zap.L().Sugar().Errorf("查询node节点资源信息失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nodeMetrics, model.NoErr.Msg))

}

// 查询命名空间下资源使用情况
func MetricsNamespace(c *gin.Context) {
	namespace := c.Param("namespace")

	mc, err := metrics.NewForConfig(config.GetK8sConfig())
	if err != nil {
		panic(err)
	}

	namespaceMetrics, err := mc.MetricsV1beta1().PodMetricses(namespace).List(metav1.ListOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询namespace节点资源信息失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, namespaceMetrics, model.NoErr.Msg))
}

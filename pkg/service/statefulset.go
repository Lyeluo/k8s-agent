package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	myutil "k8s.agent/pkg/util"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 根据namespace查询StatefulSets的列表
func StatefulsetList(c *gin.Context) {
	namespace := c.Param("namespace")

	listOptions, err := myutil.GetListOptions(c)
	if err != nil {
		zap.L().Sugar().Errorf("查询Statefulset失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	statefulsetList, err := config.GetK8sClient().AppsV1().StatefulSets(namespace).List(listOptions)
	if err != nil {
		zap.L().Sugar().Errorf("查询Statefulset失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, statefulsetList, model.NoErr.Msg))

}

// 根据name 与 namespace删除StatefulSets
func StatefulsetDelete(c *gin.Context) {
	namespace := c.Param("namespace")
	statefulsetName := c.Param("name")

	err := config.GetK8sClient().AppsV1().StatefulSets(namespace).Delete(statefulsetName, &metav1.DeleteOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除Statefulset失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nil, model.NoErr.Msg))
}

// 根据名称查询
func StatefulsetGet(c *gin.Context) {
	namespace := c.Param("namespace")
	statefulsetName := c.Param("name")

	statefulsetResult, err := config.GetK8sClient().AppsV1().StatefulSets(namespace).Get(statefulsetName, metav1.GetOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除Statefulset失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, statefulsetResult, model.NoErr.Msg))
}

// 创建StatefulSets
func StatefulsetCreate(c *gin.Context) {
	statefulset := appsv1.StatefulSet{}

	err := c.BindJSON(&statefulset)
	if err != nil {
		zap.L().Sugar().Errorf("创建Statefulset失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	statefulsetResult, err := config.GetK8sClient().AppsV1().StatefulSets(statefulset.Namespace).Create(&statefulset)
	if err != nil {
		zap.L().Sugar().Errorf("创建Statefulset失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, statefulsetResult, model.NoErr.Msg))
}

// 更新StatefulSets（删除原有然后重新部署）
func StatefulsetUpdate(c *gin.Context) {
	statefulset := appsv1.StatefulSet{}

	err := c.BindJSON(&statefulset)
	if err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	statefulsetResult, err := config.GetK8sClient().AppsV1().StatefulSets(statefulset.Namespace).Update(&statefulset)
	if err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, statefulsetResult, model.NoErr.Msg))

}

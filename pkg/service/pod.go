package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

//查询当前deployment下所有pod（根据label匹配）
func PodListByDeployment(c *gin.Context) {
	namespace := c.Param("namespace")

	deploymentName := c.Param("deployment")

	deployment, err := config.GetK8sClient().AppsV1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询pod失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, deployment, err.Error()))
		return
	}

	listSelector := model.GetK8sSelectorByMap(deployment.Spec.Selector.MatchLabels)
	option := metav1.ListOptions{
		LabelSelector: listSelector.String(),
	}
	podList, err := config.GetK8sClient().CoreV1().Pods(namespace).List(option)
	if err != nil {
		zap.L().Sugar().Errorf("查询pod失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, podList, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewResponse(true, podList, model.NoErr.Msg))
}

// 更改pod（patch）
func PodPatch(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	playLoadBytes, _ := json.Marshal(data)
	podResult, err := config.GetK8sClient().CoreV1().Pods(namespace).Patch(name, types.StrategicMergePatchType, playLoadBytes)
	if err != nil {
		zap.L().Sugar().Errorf("更新pod失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, podResult, model.NoErr.Msg))
}

// 查询日志
func PodLogs(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	podLogOpts := v1.PodLogOptions{}

	if err := c.BindJSON(&podLogOpts); err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	req := config.GetK8sClient().CoreV1().Pods(namespace).GetLogs(name, &podLogOpts)
	podLogs, err := req.Stream()
	if err != nil {
		zap.L().Sugar().Errorf("查询日志失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		zap.L().Sugar().Errorf("查询日志失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	str := buf.String()
	fmt.Printf("查询到日志：%s \n", str)
	c.JSON(http.StatusOK, model.NewResponse(false, str, model.NoErr.Msg))
}

func PodExec(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	podExecl := model.PodExecl{}
	if err := c.BindJSON(&podExecl); err != nil {
		zap.L().Sugar().Errorf("执行容器命令失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	ExecInPod(config.GetK8sClient(), namespace, name, podExecl.Command, podExecl.ContainerName)

}

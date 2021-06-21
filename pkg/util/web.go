package util

import (
	"github.com/gin-gonic/gin"
	"k8s.agent/pkg/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetListOptions(c *gin.Context) (metav1.ListOptions, error) {
	// 获取分页参数
	limit := c.Query("limit")
	continueStr := c.Query("continue")
	// 获取查询参数
	conditionSelector := model.ConditionSelector{}
	var listOptions metav1.ListOptions
	// c.BindJSON post参数必须有
	// c.ShouldBind post参数可以没有
	if err := c.ShouldBind(&conditionSelector); err != nil {

		return listOptions, err
	}

	listSelector := model.GetK8sListSelector(&conditionSelector)
	listOptions = metav1.ListOptions{
		Limit:         ParseInt64(limit, 0),
		Continue:      continueStr,
		FieldSelector: listSelector.String(),
	}

	return listOptions, nil
}

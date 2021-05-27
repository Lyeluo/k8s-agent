package model

import (
	"k8s.io/apimachinery/pkg/fields"
)

type ConditionSelector struct {
	EqualSelector map[string]string `JSON:"equalSelector"`

	NotEqualSelector map[string]string `JSON:"notEqualSelector"`
}

// 拼接list查询条件
func GetK8sListSelector(con *ConditionSelector) fields.Selector {

	equalSelector := make([]fields.Selector, 0, len(con.EqualSelector))
	for filedName, filedValue := range con.EqualSelector {

		s := fields.OneTermEqualSelector(filedName, filedValue)

		equalSelector = append(equalSelector, s)
	}

	notEqualSelector := make([]fields.Selector, 0, len(con.NotEqualSelector))
	for filedName, filedValue := range con.NotEqualSelector {

		s := fields.OneTermEqualSelector(filedName, filedValue)

		notEqualSelector = append(notEqualSelector, s)
	}
	allSelector := append(equalSelector, notEqualSelector...)

	return fields.AndSelectors(allSelector...)
}

// 组合一个map为selector格式
func GetK8sSelectorByMap(laberMap map[string]string) fields.Selector {

	selector := make([]fields.Selector, 0, len(laberMap))
	for filedName, filedValue := range laberMap {

		s := fields.OneTermEqualSelector(filedName, filedValue)

		selector = append(selector, s)
	}

	return fields.AndSelectors(selector...)
}

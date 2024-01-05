package dorm

import (
	"reflect"
	"strings"
)

type Model struct {
}

func (m *Model) getTableName(obj interface{}) string {
	return strings.ToLower(reflect.TypeOf(obj).Name())
}

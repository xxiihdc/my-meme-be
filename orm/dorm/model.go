package dorm

import (
	"fmt"
	"reflect"
	"strings"
)

type ModelI interface {
	GetTableName() string
}

type Model struct {
}

func (m *Model) GetTableName() string {
	fmt.Println("in model.go")
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(*m).Name())
	return strings.ToLower(reflect.TypeOf(m).Name())
}

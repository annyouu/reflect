package main

import (
	"fmt"
	"reflect"
)

func inspect(i interface{}) {
	v := reflect.ValueOf(i)

	// Kindを使って、型の種類を判別する
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("これはint型です", v.Int())
	case reflect.String:
		fmt.Println("これはstring型です", v.String())
	case reflect.Slice:
		fmt.Println("これはslice型です")
	default:
		fmt.Println("未知の型です:", v.Kind())
	}
}

func main() {
	inspect(42)
	inspect("hello")
	inspect([]int{1,2})
}
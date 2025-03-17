package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 普通の値
	var n int = 100
	v1 := reflect.ValueOf(n) // 100

	// reflect.Indirectでポインタでない値をそのまま返す
	fmt.Println(reflect.Indirect(v1))  // 100

	// ポインタ
	np := &n
	v2 := reflect.ValueOf(np)
	// reflect.Indirectでポインタが指す値を返す
	fmt.Println(reflect.Indirect(v2))
}
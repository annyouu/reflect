package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var n int = 10
	v := reflect.NewAt(reflect.TypeOf(0), unsafe.Pointer(&n))

	// 値を変更
	v.Elem().SetInt(100)

	// 元の変数にも反映される
	fmt.Println(n) // 100
}
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 関数fを定義
	f := func(n int) {
		fmt.Println(n)
	}

	// fのreflect.Valueを取得
	fv := reflect.ValueOf(f)

	// Callを使って関数fを呼び出す
	fv.Call([]reflect.Value{reflect.ValueOf(100)})
}
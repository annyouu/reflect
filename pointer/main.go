package main

import (
	"fmt"
	"reflect"
)

func checkType(val interface{}) {
	// 引数の型をリフレクションで取得
	v := reflect.ValueOf(val)

	// v.Kind()で変数の情報を取得
	if v.Kind() == reflect.Ptr {
		fmt.Println("ポインタ型です")
	} else {
		fmt.Println("ポインタ型ではないです")
	}
}


func main() {
	var x int
	checkType(&x) 

	var y string
	checkType(y)
}


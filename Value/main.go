package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 42

	// reflect.ValueOfでnumの値を取得
	v := reflect.ValueOf(num)

	// 値を取得
	fmt.Println("値:", v)

	// 値の型を取得
	fmt.Println("型:", v.Type())

	// 値をintとして取得する
	fmt.Println("intに変換:", v.Int())
}
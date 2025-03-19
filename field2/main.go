package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 構造体を作成
	a := struct{ N int }{N: 100}

	// reflect.ValueOf(&a)でポインタを取得
	v := reflect.ValueOf(&a)

	// Elem()でポインタの指す値を取得
	v.Elem().Field(0).SetInt(200)

	// 構造体の値を確認
	fmt.Println(a) // 200
}
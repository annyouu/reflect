package main

import (
	"fmt"
	"reflect"
)

// 構造体A
type A struct {
	n int
}

// 構造体B(Aをフィールドとして持つ)
type B struct {
	s string
	a A
}

func main() {
	//構造体Bの値を作成
	v := reflect.ValueOf(B{s: "hoge", a: A{n: 100}})

	// フィールド情報を取得
	fmt.Println(v.Field(0))

	fmt.Println(v.FieldByIndex([]int{1, 0}))

	fmt.Println(v.FieldByName("s"))
}
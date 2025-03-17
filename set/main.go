package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int = 100

	// reflect.ValueOf(n)を使うと、nの値がコピーされ、元の変数にアクセスできない
	v1 := reflect.ValueOf(n)
	if !v1.CanSet() {
		fmt.Println("v1にはSetできません")
	}

	// npはポインタを渡すことで、ポインタの指す先にアクセスすることができる
	np := &n
	v2 := reflect.ValueOf(np).Elem() // ポインタの指すnの値を取得する
	if v2.CanSet() {
		v2.Set(reflect.ValueOf(200))
	}
	fmt.Println(n)  // 200が表示される
}


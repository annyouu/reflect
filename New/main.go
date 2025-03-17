package main

import (
	"fmt"
	"reflect"
)

func main() {
	// int型のTypeを取得
	t := reflect.TypeOf(0)

	// tのゼロ値のポインタを作成
	v := reflect.New(t)

	// 出力
	fmt.Println(v.Kind()) // ptr (ポインタ型)
	fmt.Println(v.Elem().Kind()) // int (元の型)
	fmt.Println(v.Elem()) // 0
}
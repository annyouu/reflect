package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 定義した関数
	f := func(x int, y string) (int, error) { return x, nil }

	// 関数fのreflect.Typeを取得
	t := reflect.TypeOf(f)

	// 引数の数
	fmt.Println("Number of inputs:", t.NumIn()) // 出力2

	// 引数の型
	fmt.Println("First input type:", t.In(0)) // 出力 int
	fmt.Println("Second input type:", t.In(1)) // string

	// 戻り値の数
	fmt.Println("出力数:", t.NumOut()) 

	// 戻り値の型
	fmt.Println("First output type:", t.Out(0)) // 出力 int
	fmt.Println("Second outpur type:", t.Out(1)) // error

}
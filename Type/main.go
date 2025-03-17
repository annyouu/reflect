package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 42
	t := reflect.TypeOf(num)

	fmt.Println(t)  
	fmt.Println(t.Kind())
}


// reflect.TypeOf(num)：int型の情報を取得する
// t.Kind() → intの基本型を取得する
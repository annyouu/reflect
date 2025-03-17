package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "hello"
	v := reflect.ValueOf(str)
	t := v.Type()

	fmt.Println(t)   // string
	fmt.Println(t.Kind()) // string
}
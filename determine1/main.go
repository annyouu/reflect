package main

import (
	"fmt"
	"reflect"
)

func main() {
	// data := 100 → reflect.Int
	var data interface{} = 100
	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.Int:
		fmt.Println(v.Int()) // 結果: 100
	case reflect.String:
		fmt.Println(v.String())
	case reflect.Bool:
		fmt.Println(v.Bool())
	}

	// dataがhelloのとき、reflect.String
	data = "hello"
	v = reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println(v.Int())
	case reflect.String:
		fmt.Println(v.String())  // 結果 hello
	case reflect.Bool:
		fmt.Println(v.Bool())
	}

	// data
	data = true
	v = reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println(v.Int())
	case reflect.String:
		fmt.Println(v.String())
	case reflect.Bool:
		fmt.Println(v.Bool())  // 結果: true
	}
}
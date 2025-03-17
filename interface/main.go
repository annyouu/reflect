package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

func main() {
	t := reflect.TypeOf(User{})

	fmt.Println(t.Name()) // User
	fmt.Println(t.Kind()) // struct
	fmt.Println(t.NumField()) // 2 (フィールドの数)

	// フィールド情報を取得
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name, field.Type)
	}
}
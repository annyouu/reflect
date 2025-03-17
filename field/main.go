package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" db:"user_name"`
	Age int `json:"age"`
	Email string `json:"email,omitempyt"`
}

func main() {
	// User型の情報を取得
	u := User{}
	t := reflect.TypeOf(u)

	// フィールドごとにタグを取得
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("フィールド: %s, JSONタグ: %s, DBタグ: %s\n",
			field.Name,
			field.Tag.Get("json"),
			field.Tag.Get("db"))
	}
}
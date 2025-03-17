// reflect.Typeを使うことで、動的に型を調べたり、構造体のフィールド情報を取得できる。
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
	u := User{"Alice", 30}
	t := reflect.TypeOf(u)

	fmt.Println("Struct Name:", t.Name()) // Struct Name: User
	fmt.Println("Package Path:", t.PkgPath()) // Structの定義パッケージ

	// フィールド情報を取得
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: %s (%s)\n", i, field.Name, field.Type)
	}
}
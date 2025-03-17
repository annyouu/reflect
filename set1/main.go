package main

import (
	"errors"
	"fmt"
	"reflect"
)

func setString(val any, str string) error {
	// valがまずポインタかどうかを確認する
	v := reflect.ValueOf(val)

	// ポインタでない場合、エラーを返す
	if v.Kind() != reflect.Ptr {
		return errors.New("引数はポインタでなければなりません")
	}

	// ポインタなら、次にstring型かどうかを確認する
	if v.Elem().Kind() == reflect.String {
		// trueなら、ポインタの指す変数に文字列を代入する
		v.Elem().SetString(str)
		return nil
	}

	// stringでないなら、エラーを返す。
	return errors.New("ポインタはString型出なければなりません")
}

func main() {
	var s string
	err := setString(&s, "hoge")
	fmt.Println(s, err == nil)

	var n int
	err = setString(&n, "hoge")
	fmt.Println(n, err == nil)
}



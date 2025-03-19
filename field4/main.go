package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// nという非公開フィールドを持つ構造体を作成
	a := struct{ n int }{n: 100}

	// aのポインタを取得する
	v := reflect.ValueOf(&a)

	// フィールドnの値にアクセス
	fv1 := v.Elem().Field(0)

	// 変更可能か確認
	fmt.Println(fv1.CanSet()) // false(通常は変更不可)

	// unsafe.Pointerを使ってフィールドのアドレスを取得
	ptr := unsafe.Pointer(fv1.UnsafeAddr())

	// reflect.NewAtを使ってフィールドの値を変更できるようにする
	fv2 := reflect.NewAt(fv1.Type(), ptr)

	// 変更可能か確認し、値を変更
	if fv2.Elem().CanSet() {
		fv2.Elem().SetInt(200)
	}

	// 値を出力
	fmt.Println(a)
}
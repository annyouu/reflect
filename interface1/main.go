package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 型情報の基本となる rtype
type rtype struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	tflag      uint8
	align      uint8
	fieldAlign uint8
	kind       uint8  // 型の種類
}

// 構造体のフィールド情報
type structField struct {
	name string
}

// structType (rtypeを埋め込んで拡張)
type structType struct {
	rtype
	fields []structField // 構造体のフィールド情報
}

// rtypeのKind()メソッドを定義(reflect.Structを返す)
func (t *rtype) Kind() reflect.Kind {
	return reflect.Kind(t.kind)
}

// rtypeのNumField()を定義
func (t *rtype) NumField() int {
	// struct出ないのなら、エラーを出す
	if t.Kind() != reflect.Struct {
		panic("構造体型ではありません！")
	}
	// rtypeをstructTypeにキャストする
	tt := (*structType)(unsafe.Pointer(t))
	return len(tt.fields) // フィールド数を返す
}


func main() {
	// フィールドを持つ構造体型を作成
	st := structType{
		rtype: rtype{
			kind: uint8(reflect.Struct),
		},
		fields: []structField{
			{name: "Field1"},
			{name: "Field2"},
		},
	}

	// rtypeのポインタを取得する
	var t *rtype = &st.rtype

	// フィールド数を取得
	fmt.Println("NumField:", t.NumField())
}
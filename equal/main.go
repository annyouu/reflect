package main

import (
	"fmt"
	"reflect"
)

type A struct {
	N int
}

type B struct {
	A *A
	M int
}

func main() {
	b1 := &B{A: &A{N: 100}, M: 200}
	b2 := &B{A: &A{N: 100}, M: 200}

	// 演算子では、ポインタのアドレスが異なるため,false
	fmt.Println(b1 == b2)

	// DeepEqualでは中身を比較する
	fmt.Println(reflect.DeepEqual(b1, b2))
}
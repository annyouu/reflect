package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		Name string
		Age int
	}

	p := Person{Name: "Alice", Age: 30}
	v := reflect.ValueOf(p)

	fmt.Println("Type:", v.Type())
	fmt.Println("Kind:", v.Kind())
}
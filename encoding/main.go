package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	u := User{Name: "Alice", Age: 25}

	// 構造体をJSONに変換
	jsonData, _ := json.Marshal(u)
	fmt.Println(string(jsonData))
}
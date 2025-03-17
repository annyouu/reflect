package main

import (
	"os"
	"text/template"
)

type User struct {
	Name string
	Age int
	Email string
}

func main() {
	tmpl := template.Must(template.New("user").Parse("名前: {{.Name}}, 年齢: {{.Age}}, メール: {{.Email}}\n"))
	u := User{Name: "Charlie", Age: 40, Email: "charlie@example.com"}

	// テンプレートを標準出力に適用
	tmpl.Execute(os.Stdout, u)
}


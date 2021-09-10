package main

import (
	"fmt"
)

// グローバル定義のvar_test
var var_test = "package_var"

func main() {
	fmt.Println("Hello, Wodld!")
	print("Hello, World")
	println("Hello, World", 1, 2, 3)
	a := [3]string{
		"sugahara",
		"test",
		"test", // コンマがないと;が挿入されてコンパイルエラーになる
	}
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Println(var_test) // ここでは、グローバル定義を参照する

	// ローカル定義のvar_test
	var var_test = "local"
	Foo()
	fmt.Println(var_test)

	b := byte(255)
	b = b + 1
	fmt.Printf("%#v\n", b)

	c := 1.3 + 4.2i
	fmt.Printf("%#v\n", c)

	r := '松'
	fmt.Printf("%#v\n", r)
	fmt.Printf("%#v\n", string(r))
}

func Foo() {
	fmt.Println(var_test)
}

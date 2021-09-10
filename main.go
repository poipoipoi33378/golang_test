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

	var x interface{} = 3.9
	i, err := x.(int)
	f, err := x.(float32)

	fmt.Printf("%#v\n", err)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%#v\n", f)

	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	}
	// 配列
	var array [10]int
	fmt.Printf("%#v\n", array)
	fmt.Printf("%#v\n", array[9])

	array2 := [...]int{1, 2, 3, 4}
	fmt.Printf("%#v\n", array2)

	// スライス
	s := make([]int, 10)
	fmt.Printf("%#v\n", s)
	s = make([]int, 11)
	fmt.Printf("%#v\n", s[10])

	s0 := []int{1, 2, 3}
	s1 := []int{11, 12, 13}
	s2 := append(s0, s1...)
	fmt.Printf("%#v\n", s0)
	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s2)
}

func Foo() {
	fmt.Println(var_test)
}

# スターティングGo言語をやってみる
# ３章言語の基本
Goは「Cの高速性とスクリプト言語の手軽さ」を両立させるという目的を目指して設計されている。

## [事前宣言された識別子](https://zenn.dev/skanehira/articles/2021-01-14-go-identifiers-redeclared)
```
Types:
	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap close complex copy delete imag len
	make new panic print println real recover
```
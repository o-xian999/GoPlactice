package main

import (
	"fmt"
	"reflect" //追加
)

func main() {
	// 数値型のデータ型
	num01 := 123
	var num02 int = 1234567890
	num03 := 1.23
	var num04 float64 = 1.23456789

	fmt.Println(num01)
	fmt.Println(num02)
	fmt.Println(num03)
	fmt.Println(num04)

	// データ型を返す関数TypeOf
	fmt.Println(reflect.TypeOf(num01))
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(reflect.TypeOf(num03))
	fmt.Println(reflect.TypeOf(num04))
}
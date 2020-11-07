package main
import (
	"fmt"
	"reflect"
)

func main() {
	// 文字列のデータ型
	var string_a = "Hello, World!" // データ型表記
	string_b := "hello, World!" // データ型省略
	
	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))

}
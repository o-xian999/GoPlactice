package main

import (
	"fmt"
)

func main()  {
	num := 1
	num01 := 2
	num_01 := 3

	// 01num := 4 // error
	// num$01 := 5 // error

	// 1文字目が大文字:他のパッケージからも使える変数・関数
	Num := 6
	// 1文字目が小文字:そのパッケージだけで使える変数・関数
	nUM := 7

	fmt.Println(num)
	fmt.Println(num01)
	fmt.Println(num_01)
}
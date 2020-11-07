package main

import (
	"fmt"
)

func main()  {
	//名前の後にデータ型
	var num int

	// こんな感じで書いても、Goがデータ型を勝手に判断してくれる
	// var num

	// 変数の宣言と代入を同時に行うにはこんな感じで書く(よく使う)
	// num := 1

	num = 1
	// print line
	fmt.Println(num)
}
package main

import "fmt"

func main() {

	// 変数に関数を代入
	hello := func(greeting string) { // funcの後に関数名がない: 無名関数
		fmt.Println(greeting)
	}

	//無名関数はこんな感じでも書ける
	func(greeting string) {
		fmt.Println(greeting)
	}("Good evening") // 引数を後ろに書いている

	hello("Good morning")
}

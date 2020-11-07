package main

import (
	"fmt"
)

func main() {
	age := 22

	// 条件のところに()は書かない
	if age >= 20 {
		fmt.Println("abult")
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}

	// 条件式の中で宣言ができる
	if age2 := 16; age2 >= 20 {
		fmt.Println("abult")
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
	// fmt.Println(age2) // error: age2のスコープはif内
}
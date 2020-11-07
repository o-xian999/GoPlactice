package main

import "fmt"

// 引数なし関数
func sayHello() {
	fmt.Println("Hello, World!")
}

// 引数あり関数
func sayString(greeting string) {
	fmt.Println(greeting)
}

func main() {
	sayHello()
	sayString("Good morning")
}

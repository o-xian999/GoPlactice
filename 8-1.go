package main

import (
	"fmt"
)

func main() {
	a := [3]string {"sato", "suzuli", "takahashi"}
	// a := [...]string {"sato", "suzuli", "takahashi"} // 要素数は省略OK:"..."を書く
	a[1] = "tanaka"

	b := [2][2]string {{"sato", "suzuki"}, {"tanaka", "takahasi"}}

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(b[0][0])
	fmt.Println(b[1][1])
	fmt.Println(b[1][0])
}
package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Println(i)
	}

	i := 0
	// for {
	// 	fmt.Println(i) //無限ループ
	// }

	for i <= 4 {
		fmt.Println(i)
		i++
	}

}

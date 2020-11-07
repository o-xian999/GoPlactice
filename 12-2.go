package main

import "fmt"

func cal(x, y int) { // (x int, y int)と同じ意味
	fmt.Println(x / y)
}

// 引数2、返り値1
func cal2(x, y int) int {
	return (x / y)
}

// 引数2、返り値2
func cal3(x, y int) (int, int) {
	return (x / y), (x * y)
}

// 中で変数使う場合
func cal5(x, y int) (int, int) {
	a := x / y
	b := x * y
	return a, b
}

// 戻り地2つあるとこんなふうにも書ける
func cal4(x, y int) (a int, b int) {
	a = x / y
	b = x * y
	return // 戻り値を宣言しているので、書かなくても動く
	// return a, b
}

func main() {
	cal(6, 4)
	fmt.Println(cal2(10, 5))

	// こういう書き方ができる
	result1, result2 := cal3(20, 4)
	fmt.Println(result1, result2) // 5 80

	fmt.Println(cal4(60, 20)) // 3 1200
}

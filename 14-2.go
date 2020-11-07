package main

import "fmt"

type Student struct {
	name string
}

// メソッド(構造体などの特定の型に関連付けられた関数)を定義
// 引数ありメソッド
func (s Student) avg(math, english float64) {
	fmt.Println((math + english) / 2)
}

// 返り値ありメソッド
func (s Student) avg2(math, english float64) float64 {
	return (math + english) / 2
}

// 返り値前もって宣言するメソッド
func (s Student) avg3(math, english float64) (avgResult float64) {
	avgResult = (math + english) / 2
	return
}

func main() {
	a001 := Student{"sato"}
	a001.avg(80, 70)
	fmt.Println(a001.avg2(80, 70))
	fmt.Println(a001.avg3(80, 70))
}

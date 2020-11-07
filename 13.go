package main

import "fmt"

type Student struct { // 構造体の定義
	name          string
	math, english float64 // フィールドを定義
}

func main() {
	var s Student // 構造体の初期化
	s.name = "sato"
	s.math = 80
	s.english = 70

	t := Student{"sato", 80, 70} // これでもいける
	// r := Student{name: "sato", english: 70, math: 80} // 順番変えるならこんな感じ

	fmt.Println(s.name)
	fmt.Println(t.math)
}

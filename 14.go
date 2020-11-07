package main

import "fmt"

type Student struct {
	name          string
	math, english float64
}

 // メソッド(構造体などの特定の型に関連付けられた関数)を定義
func (s Student) avg() {　// メソッドのfunc直後の()をレシーバという
	fmt.Println(s.name, (s.math+s.english)/2)
}

func main() {
	a001 := Student{"sato", 80, 70}
	a001.avg() // メソッドはこんな感じで使うんやで
}

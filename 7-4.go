package main
import (
	"fmt"
	"reflect"
)

func main() {
	// boolのデータ型(Boolenだな)
	a := 10
	b := 1
	// bool型の変数num_bool
	var num_bool bool = a > b

	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))

}
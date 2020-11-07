// 最低1つのpackageに所属している必要がある
package main
// fmtをimport
import ("fmt")

// mainパッケージ内のmain関数からスタートする
// プログラム開始地点：エントリーポイント
func main()  { // プログラム開始地点：エントリーポイント

	//fmtパッケージ内のPrintln関数を使って、文字列を表示
	fmt.Println("Good morning")

	fmt.Println("Good afternoon")
	
	fmt.Println("Good evening")
}

//コンパイルのみ(実行ファイル生成): go build <filename>
//コンパイルと実行を同時: go run <filename>
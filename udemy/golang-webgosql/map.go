package main

import "fmt"

func main() {
    m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	/*
	banana 200
    apple 100
	*/

	// 要素追加
	fmt.Println(m["apple"]) // 100
	m["grape"] = 400
	fmt.Println(m) // map[apple:100 banana:200 grape:400]

	// エラーハンドリング
	_, ok := m["hoge"] // 変数okにはboolが入る。hogeというキーは存在しないのでfalseが代入
	if !ok {
		fmt.Println("error")
	}
	// error
}

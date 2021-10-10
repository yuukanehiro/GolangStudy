package main

import "fmt"

func main() {
	sl := []int{100, 200}
	fmt.Println(sl) // [100 200]

	sl = append(sl, 300)
	fmt.Println(sl) // [100 200 300]

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl) // [100 200 300 400 500 600]

	sl2 := make([]int, 5) // 配列の作成
	fmt.Println(sl2) // [0 0 0 0 0]
	fmt.Println(len(sl2)) // 5 ... 要素数取得
	fmt.Println(cap(sl2)) // 5 ... 容量 ... メモリを気にするような時に使う

	sl3 := make([]int, 5, 10) // ... make([]T, 初期値で埋める要素数, 容量)
	sl4 := append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl4)) // 12
	fmt.Println(cap(sl4)) // 20
}
package main

import "fmt"

func main() {
	// make(), apped(), copy()
	slice1 := make([]int, 3) 
	fmt.Println(slice1) // [0 0 0]

	// 短縮系でのスライス定義
	slice2 := []int{1, 3, 5}
	fmt.Println(slice2) // [1 3 5]
	// append() スライスに要素の追加
	fmt.Println(append(slice2, 6, 8, 10)) // [1 3 5 6 8 10]
	// copy
	copy(slice1, slice2) // slice1にslice2を複製
	fmt.Println(slice1) // [1 3 5]
}
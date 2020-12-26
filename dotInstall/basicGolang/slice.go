package main

import "fmt"

func main() {
	array1 := [5]int{1, 3, 5, 7, 11}
	slice1 := array1[2:4]
	slice2 := array1[2:]
	fmt.Println(slice1)    // [5 7]
	fmt.Println(slice1[1]) // 7
	fmt.Println(slice2)    // [5 7 11]
	// 要素数を調べる
	fmt.Println(len(slice2)) // 3
	// スライスの最大容量を示す
	fmt.Println(cap(slice2)) // 3
}
package main

import "fmt"

func main() {
	var array1 [6]int // array1[0] - array1[5]
	array1[2] = 3
	array1[4] = 20
	fmt.Println(array1[2]) // 3

	array2 := [...]int{1, 3, 5}
	fmt.Println(array2[2]) // 5
}
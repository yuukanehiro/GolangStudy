package main

import "fmt"

func main() {
	array1 := [5]int{1, 3, 5, 7, 11}
	slice1 := array1[2:4]
	fmt.Println(slice1) // [5 7]
}
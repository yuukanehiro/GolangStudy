package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
    fmt.Println(Sum(1, 2, 10)) // 13

	sl := []int{4, 5}
	fmt.Println(Sum(sl...)) // 9
}
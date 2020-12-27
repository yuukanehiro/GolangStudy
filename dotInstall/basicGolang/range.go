package main

import "fmt"

func main() {
	slice1 := []int{1, 3, 5, 7}
	for i, v := range slice1 {
		fmt.Println(i, v)
							// 0 1
							// 1 3
							// 2 5
							// 3 7
	}
	for _, v := range slice1 {
		fmt.Println(v)
							// 1
							// 3
							// 5
							// 7
	}

	map1 := map[string]int{"yuu":100, "tanaka": 500}
	for k, v := range map1 {
		fmt.Println(k, v)
							// yuu 100
							// tanaka 500
	}
}
package main

import "fmt"

func main() {
    var sl []int
	fmt.Println(sl) // []

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2) // [100 200]

	sl3 := []string{"A", "B"}
	fmt.Println(sl3) // [A B]

	sl4 := make([]int, 5)
	fmt.Println(sl4) // [0 0 0 0 0]

	sl2[0] = 999
	fmt.Println(sl2) // [999 200]

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5) // [1 2 3 4 5]
	fmt.Println(sl5[0]) // 1
	fmt.Println(sl5[2:4]) // [3 4]
	fmt.Println(sl5[:2]) // [1 2]
	fmt.Println(sl5[2:]) // [3 4 5]
	fmt.Println(sl5[:]) // [1 2 3 4 5]
	fmt.Println(sl5[len(sl5)-1]) // 最後の要素を取り出す ... 5
	fmt.Println(sl5[1:len(sl5)-1]) // [2 3 4]
}
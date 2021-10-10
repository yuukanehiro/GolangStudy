package main

import "fmt"

func main() {
	sla := []string{"PHP", "GO"}
	slb := sla
	slb[0] = "Kotlin"
	fmt.Println(sla) // [Kotlin GO] ... PHPが上書きされてしまった

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	n:= copy(sl2, sl) // sl2にslをコピーして代入。nにはコピーできた要素数が代入される
	fmt.Println(n, sl2) // 5 [1 2 3 4 5]
}
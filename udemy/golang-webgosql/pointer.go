package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n) // 100
	fmt.Println(&n) // 0xc00012a008

	Double(n)
	fmt.Println(n) // 100

	var p *int = &n // &変数 ... ポインタ
	fmt.Println(p) // 0xc000132008
	fmt.Println(*p) // 100 ... *をつけると実体

	*p = 300 // *pで実体を指定して代入
	fmt.Println(n) // 300
	n = 200
	fmt.Println(*p) // 200

	Doublev2(&n)
	fmt.Println(n) // 400

	// スライスは参照型なので関数先で自動的に値が参照渡しされてる
	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl) // [2 4 6]
}
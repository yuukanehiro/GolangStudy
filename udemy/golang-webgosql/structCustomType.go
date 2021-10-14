package main

import "fmt"

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}


func main() {
	var mi MyInt
	fmt.Println(mi) // 0
	fmt.Printf("%T\n", mi) // main.MyInt

	mi.Print() // 0
}
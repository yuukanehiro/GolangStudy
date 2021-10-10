package main

import "fmt"


func main() {
	var i2 int64 = 200
	fmt.Println(i2 + 50) // 250

	fmt.Printf("%T\n", i2) // int64

	fmt.Printf("%T\n", int32(i2)) // int32
}
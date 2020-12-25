package main

import "fmt"

func main() {
	number := 12
	var numberPointer * int
	numberPointer = &number

	fmt.Println(numberPointer) // 0xc000016070
	fmt.Println(*numberPointer) // 12
}
package main

import "fmt"

func main() {
	i := interface{}("hello1")
	s := i.(string)
	fmt.Println(s) // hello1

	// i2 := interface{}("hello2")
	// s2 := i2.(int)
	// fmt.Println(s2) // panic: interface conversion: interface {} is string, not int
}
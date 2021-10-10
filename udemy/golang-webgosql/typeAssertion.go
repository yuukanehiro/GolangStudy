package main

import (
	"fmt"
)

func main() {
	i := interface{}("hello1")
	s := i.(string)
	fmt.Println(s) // hello1

	// i2 := interface{}("hello2")
	// s2 := i2.(int)
	// fmt.Println(s2) // panic: interface conversion: interface {} is string, not int

	i3 := interface{}("hello3")
    n, ok := i3.(int) // 引数を2つ取ることでpanicを防止
    fmt.Println(n, ok) // 0 false ... panicにならない

    s3, ok := i3.(string)
    fmt.Println(s3, ok) // hello3 true

	var x interface{} = 3
	if x == nil {
		fmt.Println("None")
	} else if _, isInt := x.(int); isInt {
		fmt.Println(x, "x is Int")
	} else if s4, isString := x.(string); isString {
		fmt.Println(s4, "x is String")
	} else {
		fmt.Println("I don't know")
	}
	// 3 x is Int

	// switchで再現
	switch x.(type) {
	case int:
		fmt.Println(x, "x is Int")
	case string:
		fmt.Println(x, "x is String")
	default:
		fmt.Println("I don't know")
	}
	// 3 x is Int

	// switchで再現
	switch v5 := x.(type) {
	case int:
		fmt.Println(v5, "v5 is Int")
	case string:
		fmt.Println(v5, "v5 is String")
	default:
		fmt.Println("I don't know")
	}
	// 3 v5 is Int
}
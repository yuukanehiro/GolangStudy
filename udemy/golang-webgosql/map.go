package main

import "fmt"

func main() {
    m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	/*
	banana 200
    apple 100
	*/
}

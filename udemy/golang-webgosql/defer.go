package main

import "fmt"

func TestDefer() {
	defer fmt.Println("End")

	fmt.Println("Start")
}

func main() {
	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	// Start
	// End
	// 1
	// 2
	// 3
}
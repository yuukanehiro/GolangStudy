package main

import "fmt"

func main() {
	const message = "yuuです"
	fmt.Println(message) // yuuです

	const (
		sun = iota
		mon
		tue
	)
	println(sun, mon, tue) // 0 1 2
}





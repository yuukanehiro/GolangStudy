package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "A"
	ch1 <- 1
	ch2 <- "B"
	ch1 <- 2

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	}
	// ランダムになる
	// kanehiroyuu@MacBook-Pro-2 golang-webgosql % go run channelSelect.go
	// A!?
	// kanehiroyuu@MacBook-Pro-2 golang-webgosql % go run channelSelect.go
	// A!?
	// kanehiroyuu@MacBook-Pro-2 golang-webgosql % go run channelSelect.go
	// 1001
	// kanehiroyuu@MacBook-Pro-2 golang-webgosql % go run channelSelect.go
	// 1001
	// kanehiroyuu@MacBook-Pro-2 golang-webgosql % go run channelSelect.go
	// A!?
}
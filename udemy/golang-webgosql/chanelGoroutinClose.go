package main

import (
	"fmt"
	"time"
)

func reciver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name, "End")
}

func main() {
    ch1 := make(chan int, 2)
	ch1 <- 1

	go reciver("1.goroutin", ch1)
	go reciver("2.goroutin", ch1)
	go reciver("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)

	// 3.goroutin 1
	// 2.goroutin 1
	// 1.goroutin 0
	// 1.goroutin 4
	// 1.goroutin 5
	// 1.goroutin 6
	// 2.goroutin 3

	// (略)

	// 1.goroutin 83
	// 1.goroutin 84
	// 1.goroutin 85
	// 1.goroutin 86
	// 1.goroutin 87
	// 1.goroutin 88
	// 2.goroutin 69
	// 3.goroutin 82
	// 3.goroutin 90
	// 3.goroutin 92
	// 3.goroutin 93
	// 3.goroutin 94
	// 1.goroutin 89
	// 1.goroutin 96
	// 2.goroutin 91
	// 2.goroutin 98
	// 2.goroutin 99
	// End 2.goroutin
	// 1.goroutin 97
	// End 1.goroutin
	// 3.goroutin 95
	// End 3.goroutin
}


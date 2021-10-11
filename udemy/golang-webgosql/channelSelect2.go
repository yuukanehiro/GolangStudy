package main

import "fmt"

func main() {
	ch3 := make(chan int,)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// reciever
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 -1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("recieved", i3)
	    default:
		    if n > 100 {
			    break L
			}
		}
	}
}

// kanehiroyuu@MacBook-Pro-2 golang-webgosql % go run channelSelect2.go
// recieved -1
// recieved 1
// recieved 3
// recieved 5
// recieved 7
// recieved 9
// recieved 11
// recieved 13
// recieved 15
// ...
// recieved 183
// recieved 185
// recieved 187
// recieved 189
// recieved 191
// recieved 193
// recieved 195
// recieved 197
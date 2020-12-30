package main

import (
	"fmt"
	"time"
)

func task1() {
	time.Sleep(time.Second * 2) // 2秒処理待ち
	fmt.Println("task1 finished!")
}

func task2() {
	fmt.Println("task2 finished!")
}

func main() {
	go task1()
	go task2()

	time.Sleep(time.Second * 3)
}
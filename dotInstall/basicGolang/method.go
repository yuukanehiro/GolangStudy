package main

import "fmt"

type user struct {
	name string
	score int
}

func (u user) show() {
	fmt.Printf("name:%s, socre:%d\n", u.name, u.score)
}

func (u user) scoreUp() {
	u.score++
}

func main() {
	userInstance := user{"tanaka", 100}
	userInstance.scoreUp()
	userInstance.show() // name:tanaka, socre:100
}
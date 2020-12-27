package main

import "fmt"

type user struct {
	name string
	score int
}

func (u user) show() {
	fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

func (u *user) scoreUp() {
	u.score++
}

func main() {
	userInstance := user{name:"tanaka", score:100}
	userInstance.scoreUp()
	userInstance.show() // name:tanaka, score:101
}
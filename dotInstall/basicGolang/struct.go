package main

import "fmt"

type user struct {
	name string
	score int
}

func main() {
	user1 := new(user)
	fmt.Println(user1) // &{ 0} ... ポインタが返却される

	user1.name = "Yuu"
	user1.score = 99
	fmt.Println(user1) // &{Yuu 99}

	user2 := user{"tanaka", 100}
	fmt.Println(user2) // {tanaka 100} ... &が消えて、値が返却される
}
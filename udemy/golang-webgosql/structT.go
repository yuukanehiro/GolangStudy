package main

import "fmt"

type User struct {
	Name string
	Age int
	// X, Y int
}

type T struct {
	User User
}

type T2 struct {
	User
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t) // {{user1 10}}
	fmt.Println(t.User) // {user1 10}
	fmt.Println(t.User.Name) // user1

	t2 := T2{User: User{Name: "user1", Age: 10}}
	fmt.Println(t2.User.Name) // user1
}
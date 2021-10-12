package main

import "fmt"

type User struct {
	Name string
	Age int
	// X, Y int
}

func (u User) sayName() {
	fmt.Println(u.Name)
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 999)
	fmt.Println(user1) // &{user1 999}
}



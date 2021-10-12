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

// これだとコピーしたインスタンスのNameを更新していることになるので、更新できない
func (u User) setName(name string) {
	u.Name = name
}

// 🐱こう書く
func (u *User) setName2(name string) {
	u.Name = name
}

func main() {
	user1 := &User{"user1", 999}
	fmt.Println(user1) // &{user1 999}
	user1.sayName() // user1

	user1.setName("A")
	user1.sayName() // user1 ... 変わってない
	user1.setName2("B")
	user1.sayName() // B ... 変わってる🐱
}
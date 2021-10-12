package main

import "fmt"

type User struct {
	Name string
	Age int
	// X, Y int
}

// コピーに対して更新していて更新できない関数
func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// こう書く🐱
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1) // { 0}
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1) // {user1 10}

	user2 := User{}
	fmt.Println(user2) // { 0}
	user2.Name = "user2"
	fmt.Println(user2) // {user2 0}

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3) // {user3 30}

	user4 := User{Name: "user4", Age: 40}
	fmt.Println(user4) // {user4 40}

	user7 := new(User)
	fmt.Println(user7) // &{ 0} ... ポインタ型になる

	// 🐱基本はこの書き方
	user8 := &User{}
	fmt.Println(user8) // &{ 0} ... ポインタ型になる

	UpdateUser(user1)
	fmt.Println(user1) // {user1 10} ... 変わらない
	UpdateUser2(user8)
	fmt.Println(user8) // &{A 1000} ... 更新することができている🐱
}
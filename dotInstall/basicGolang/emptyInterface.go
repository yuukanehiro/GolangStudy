package main

import "fmt"

func show(t interface{}) {
	// 型アサーション
	_, ok := t.(japanese)
	if ok {
		fmt.Println("日本人じゃぞ！")
	} else {
		fmt.Println("外国人じゃぞ！")
	}
}


type greeter interface {
	greet()
}

type japanese struct {}
type american struct {}

func (j japanese) greet() {
	fmt.Println("こんちゃ")
}

func (a american) greet() {
	fmt.Println("Hey!")
}


func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}

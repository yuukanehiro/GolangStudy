package main

import "fmt"

func hello(name string) {
	fmt.Println(name + "さんです")
}

func message() string {
	message := "馬鹿な、ありえぬ。"
	return message
}

func main() {
	myName := "優"
	hello(myName) // 優さんです
	fmt.Println(message()) // 馬鹿な、ありえぬ。
}
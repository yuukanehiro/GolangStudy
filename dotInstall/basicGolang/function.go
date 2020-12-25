package main

import "fmt"

func hello(name string) {
	fmt.Println(name + "さんです")
}

func main() {
	myName := "優"
	hello(myName) // 優さんです
}
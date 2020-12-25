package main

import "fmt"

func hello(name string) {
	fmt.Println(name + "さんです")
}

func message() string {
	message := "馬鹿な、ありえぬ。"
	return message
}

func manyReturn(word1 string, number1 int) (string, int) {
	return word1, number1
}

func main() {
	myName := "優"
	hello(myName) // 優さんです
	fmt.Println(message()) // 馬鹿な、ありえぬ。
	word := "damage:"
	number := 99
	fmt.Println(manyReturn(word, number)) // damage: 99
}
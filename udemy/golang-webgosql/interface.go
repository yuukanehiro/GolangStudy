package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Stringfy interface {
	ToString() string
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 10},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
	// Name=Taro, Age=10
	// Name=123-456, Age=AB-1234
}




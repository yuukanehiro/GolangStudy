package main

import "fmt"

func main() {
    i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		/*
        1
        2
        4
        5
        6
        7
        8
        9
		*/
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
		/*
		0
		1
		2
		*/
	}

	sl := []string{"Python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	/*
    0 Python
    1 PHP
    2 GO
	*/
}
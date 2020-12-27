package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["yuu"] = 20
	map1["satou"] = 50
	fmt.Println(len(map1)) // 2
	fmt.Println(map1)      // map[satou:50 yuu:20]

	delete(map1, "satou")
	fmt.Println(map1)      // map[yuu:20]
	v, ok := map1["yuu"]
	fmt.Println(v)         // 20
	fmt.Println(ok)        // true　... 存在するという意
}
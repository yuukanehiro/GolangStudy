package main

import "fmt"

func outer() {
    var s4 string = "outer"
    fmt.Println(s4)
}

func main() {
    var i int = 100
    fmt.Println(i)
    var s string = "Hello Go"
    fmt.Println(s)
    // Hello Go

    var t, f bool = true, false
    fmt.Println(t, f)

    var (
        i2 int = 200
        s2 string = "Golang"
    )
    fmt.Println(i2, s2) 
    // 200
    // Golang

    var i3 int
    var s3 string
    fmt.Println(i3, s3) // 0 と 表示されないけど空文字

    i3 = 300
    s3 = "Go"
    fmt.Println(i3, s3)

    i = 150
    fmt.Println(i)

    // 暗黙的な定義 ... 型を動的に見てくれる ... しかし可能な限り型指定する
    // 変数名 := 値
    i4 := 400
    fmt.Println(i4)

    // i4 = "Yuu3"
    // fmt.Println(i4) intで前のセクションで定義されているのでエラーになる🐱💦

    outer()
}
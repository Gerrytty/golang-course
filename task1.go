package main

import "fmt"

func main() {
    var a, first, second, third int
    fmt.Scan(&a)

    first = a / 100
    second = (a % 100) / 10
    third = a % 10

    if first != second && first != third && second != third {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }

    fmt.Println("ok")
}
package main
import "fmt"

func main() {
    var n int
    var s []int16

    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var a int16
        fmt.Scan(&a)
        s = append(s, a)
    }

    fmt.Println(s[3])
}
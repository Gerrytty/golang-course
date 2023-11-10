package main

import "fmt"

func main() {
	var a, b, res int16

	fmt.Scan(&a, &b)

	res = ((2 * a + (b-a)) * (b-a+1)) / 2

	fmt.Println(res)
}
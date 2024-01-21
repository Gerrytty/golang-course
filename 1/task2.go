package main

import "fmt"

func main() {

	var a int16
	fmt.Scan(&a)

	// reverse three-digit number
	fmt.Println((a % 10) * 100 + (a % 100 / 10) * 10 + (a / 100))
}
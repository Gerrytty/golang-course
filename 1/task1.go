package main

import "fmt"

func main() {

	var a int16
	fmt.Scan(&a)

	// sum of digits in three-digit numbers
	fmt.Println((a / 100) + (a % 10) + (a % 100 / 10))
}
package main

import "fmt"

func main() {
	var a, firstPair, secondPair int32
	fmt.Scan(&a)

	firstPair = a / 1000
	secondPair = a % 1000

	if firstPair / 100 + firstPair % 10 + firstPair / 10 % 10 == secondPair / 100 + secondPair % 10 + secondPair / 10 % 10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
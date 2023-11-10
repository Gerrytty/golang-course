package main

import "fmt"

func main() {
	var a, max, c int32
	max = -100
	c = 1
	a = -1

	for ; a != 0; fmt.Scan(&a) {
		if a > max {
			max = a
			c = 1
		} else if a == max {
			c++
		}
	}

	fmt.Println(c)
}
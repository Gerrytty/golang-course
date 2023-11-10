package main

import "fmt"

func main() {
	var n, a, c int
	fmt.Scan(&n)

	sum := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		savedA := a
		if a % 8 == 0 {
			c = 0
			for a > 0 {
				a /= 10
				c++
			}
			if c == 2 {
				sum += savedA
			}
		}
	}

	fmt.Println(sum)
}
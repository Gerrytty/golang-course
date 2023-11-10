package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	maxA := 1
	maxB := 1

	tempA := a
	tempB := b

	for tempA > 10 {
		maxA *= 10
		tempA /= 10
	}

	for tempB > 10 {
		maxB *= 10
		tempB /= 10
	}


	for a > 0 {
		num := a / maxA
		newB := b
		for newB > 0 {
			if newB % 10 == num {
				fmt.Print(num, " ")
			}
			newB /= 10
		}
		a -= num * maxA
		maxA /= 10
	}
}
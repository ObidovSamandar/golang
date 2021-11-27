package problemsolving

import (
	"fmt"
)

func CalculateFibonacciSeries(length int) {

	var x, y, result int
	fibonacciSeries := []int{}

	x = 0
	y = 1

	for result < length {
		result = x + y
		if result > length {
			break
		}
		fibonacciSeries = append(fibonacciSeries, result)
		y = x      //2
		x = result //3
	}

	fmt.Println(fibonacciSeries)
}

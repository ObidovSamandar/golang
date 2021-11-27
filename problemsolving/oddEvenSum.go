package problemsolving

import "fmt"

func OddEvenSum(arr []int) {
	var evenSum, oddSum int

	for i := 1; i < len(arr); i++ {
		if i%2 == 0 {
			evenSum += arr[i]
		} else {
			oddSum += arr[i]
		}
	}

	if evenSum == oddSum {
		fmt.Println("Yes. Sum is =", evenSum)
	} else {
		fmt.Println("No. Difference =", evenSum-oddSum)
	}
}

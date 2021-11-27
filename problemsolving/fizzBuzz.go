package problemsolving

import "fmt"

func FizzBuzz(length int) {
	fizzBuzzList := make(map[int]string)

	for i := 1; i <= length; i++ {
		if i%3 == 0 && i%5 == 0 {
			fizzBuzzList[i] = "Fizzbuzz"
		} else if i%3 == 0 {
			fizzBuzzList[i] = "Fizz"
		} else if i%5 == 0 {
			fizzBuzzList[i] = "Buzz"
		}
	}
	fmt.Println(fizzBuzzList)
}

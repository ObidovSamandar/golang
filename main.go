package main

import (
	"github.com/obidovsamandar/first_project/gopackages"
	"github.com/obidovsamandar/first_project/problemsolving"
)

func main() {
	problemsolving.CalculateFibonacciSeries(100)
	problemsolving.FizzBuzz(100)
	problemsolving.IStrPalindrom("aziza")
	problemsolving.RemoveIntDuplicatesFromArr([]int{11, 33, 41, 11, 32, 32, 3})
	problemsolving.RemoveStrDuplicatesFromArr([]string{"Tashkent", "Uzbekistan", "New York", "Mumbai", "Abu Dahli", "Tashkent", "New York"})
	problemsolving.OddEvenSum([]int{4, 10, 50, 60, 20})
	problemsolving.OddEvenSum([]int{3, 5, 8, 1})

	gopackages.ColorPackage()
}

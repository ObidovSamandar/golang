package problemsolving

import (
	"fmt"
)

func RemoveStrDuplicatesFromArr(arr []string) {
	allKeys := make(map[string]bool)
	list := []string{}

	for _, item := range arr {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	fmt.Println(list)
}

func RemoveIntDuplicatesFromArr(arr []int) {
	allKeys := make(map[int]bool)
	list := []int{}

	for _, item := range arr {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	fmt.Println(list)
}

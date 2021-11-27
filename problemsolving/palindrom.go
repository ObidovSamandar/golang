package problemsolving

import "fmt"

func IStrPalindrom(str string) {
	isPalindrom := false

	// aziza
	middleOfStrLength := int(len(str) / 2)

	for i := 0; i <= middleOfStrLength; i++ {
		if str[i] == str[len(str)-i-1] {
			isPalindrom = true
		} else {
			isPalindrom = false
			break
		}
	}

	if isPalindrom {
		fmt.Println(str, "is palindrom")
	} else {
		fmt.Println(str, "is not palindrom")
	}
}

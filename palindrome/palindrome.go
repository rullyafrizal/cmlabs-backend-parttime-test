package palindrome

import (
	"strings"
)

func IsPalindrome(v string) bool {
	// Split into array to remove whitespace in multiple word
	// Join it again, and lower the case
	vArr := strings.Split(v, " ")
	vJoined := strings.Join(vArr, "")
	vLower := strings.ToLower(vJoined)

	// loop and compare all the way to the middle
	// true palindrome will not enter the if block
	// oppositeIndex is the opposite of "i" from the middle index
	for i := 0; i < len(vLower)/2; i++ {
		var oppositeIndex = len(vLower) - i - 1

		if string(vLower[i]) != string(vLower[oppositeIndex]) {
			return false
		}
	}

	return true
}

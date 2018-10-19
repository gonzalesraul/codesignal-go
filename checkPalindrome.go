package codesignal

import (
	"regexp"
)

/*
CheckPalindrome : true if inputString is a palindrome, false otherwise.

Given the string, check if it is a palindrome (A palindrome is a string that reads the same left-to-right and right-to-left).

Example

    For inputString = "aabaa", the output should be
    checkPalindrome(inputString) = true;
    For inputString = "abac", the output should be
    checkPalindrome(inputString) = false;
    For inputString = "a", the output should be
    checkPalindrome(inputString) = true.

Input/Output

- [execution time limit] 4 seconds (go)

- [input] string inputString

    A non-empty string consisting of lowercase characters.

    Guaranteed constraints:
    1 ≤ inputString.length ≤ 10^5.

- [output] boolean
    true if inputString is a palindrome, false otherwise.
*/
func CheckPalindrome(inputString string) bool {
	len := len(inputString)
	if len == 0 || len > 100000 || regexp.MustCompile("[A-Z]").MatchString(inputString) {
		return false
	}

	for index, max := 0, len-(len%2)/2; index < max; index++ {
		len--
		if inputString[index] != inputString[len] {
			return false
		}
	}
	return true
}

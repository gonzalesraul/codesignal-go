package codesignal

/*
CommonCharacterCount - Given two strings, find the number of common characters between them.

Example

For s1 = "aabcc" and s2 = "adcaa", the output should be
commonCharacterCount(s1, s2) = 3.

Strings have 3 common characters - 2 "a"s and 1 "c".

Input/Output

- [execution time limit] 4 seconds (go)

- [input] string s1

    A string consisting of lowercase English letters.

    Guaranteed constraints:
    1 ≤ s1.length ≤ 15.

- [input] string s2

    A string consisting of lowercase English letters.

    Guaranteed constraints:
    1 ≤ s2.length ≤ 15.

- [output] integer
	Given two strings, find the number of common characters between them.
*/
func CommonCharacterCount(s1 string, s2 string) int {
	if len(s1) < 1 || len(s1) > 15 || len(s2) < 1 || len(s2) > 15 {
		return 0
	}

	if len(s1) > len(s2) { //iterate over the minor length
		s1, s2 = s2, s1
	}

	var result int
	chars := []rune(s2)
	for _, char := range s1 {
		for index := 0; index < len(chars); index++ {
			if char == chars[index] {
				result++
				chars = append(chars[:index], chars[index+1:]...)
				break
			}
		}
	}
	return result
}

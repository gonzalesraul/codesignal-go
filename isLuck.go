package codesignal

import (
	"strconv"
)

/*
IsLucky - true if n is a lucky ticket number, false otherwise.

Ticket numbers usually consist of an even number of digits.
A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.

Given a ticket number n, determine if it's lucky or not.

Example

    For n = 1230, the output should be
    isLucky(n) = true;
    For n = 239017, the output should be
    isLucky(n) = false.

Input/Output

- [execution time limit] 4 seconds (go)

- [input] integer n

    A ticket number represented as a positive integer with an even number of digits.

    Guaranteed constraints:
    10 ≤ n < 10^6.

- [output] boolean
	true if n is a lucky ticket number, false otherwise.
*/
func IsLucky(n int) bool {
	chars := strconv.Itoa(n)
	count := len(chars)
	if n < 10 || n > 999999 || count%2 == 1 {
		return false
	}

	var part1, part2 int

	for _, n := range chars[:count/2] {
		part1 += int(n)
	}
	for _, n := range chars[count/2:] {
		part2 += int(n)
	}
	return part1 == part2
}

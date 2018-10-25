package codesignal

/*
AllLongestStrings : Array of the longest strings, stored in the same order as in the inputArray.
Given an array of strings, return another array containing all of its longest strings.

Example

For inputArray = ["aba", "aa", "ad", "vcd", "aba"], the output should be
allLongestStrings(inputArray) = ["aba", "vcd", "aba"].

Input/Output

- execution time limit] 4 seconds (go)

- input] array.string inputArray

    A non-empty array.

    Guaranteed constraints:
    1 ≤ inputArray.length ≤ 10,
    1 ≤ inputArray[i].length ≤ 10.

- [output] array.string
    Array of the longest strings, stored in the same order as in the inputArray.

*/
func AllLongestStrings(inputArray []string) []string {
	if count := len(inputArray); count < 1 || count > 10 {
		return nil
	} else if count == 1 {
		return inputArray
	}

	higherLength := len(inputArray[0])
	for current := 0; current < len(inputArray)-1; {
		if len(inputArray[current]) > higherLength { // Whenever find a bigger value, cuts the slice starting from that index
			higherLength = len(inputArray[current])
			inputArray = inputArray[current:]
		}
		if next := current + 1; len(inputArray[current]) > len(inputArray[next]) {
			inputArray = append(inputArray[:next], inputArray[next+1:]...)
		} else if len(inputArray[current]) < len(inputArray[next]) {
			inputArray = append(inputArray[:current], inputArray[next:]...)
		} else {
			current++
		}
	}
	return inputArray
}

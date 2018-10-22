package codesignal

/*
AdjacentElementsProduct : The largest product of adjacent elements.

Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.

Example

For inputArray = [3, 6, -2, -5, 7, 3], the output should be
adjacentElementsProduct(inputArray) = 21.

7 and 3 produce the largest product.

Input/Output

- [execution time limit] 4 seconds (go)

- [input] array.integer inputArray
    An array of integers containing at least two elements.

    Guaranteed constraints:
    2 ≤ inputArray.length ≤ 10,
    -1000 ≤ inputArray[i] ≤ 1000.

- [output] integer
	The largest product of adjacent elements.

*/
func AdjacentElementsProduct(inputArray []int) int {

	if len(inputArray) < 2 {
		return 0
	}

	result := inputArray[0] * inputArray[1] //Guaranteed constraints (length > 2)

	for index, count := 2, len(inputArray); index < count; index++ {
		if product := inputArray[index-1] * inputArray[index]; product > result {
			result = product
		}
	}

	return result
}

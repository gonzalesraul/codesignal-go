package codesignal

/*
Add : The sum of the two inputs.

Write a function that returns the sum of two numbers.

Example

For param1 = 1 and param2 = 2, the output should be
add(param1, param2) = 3.

Input/Output

- [execution time limit] 4 seconds (go)

- [input] integer param1

    Guaranteed constraints:
    -1000 ≤ param1 ≤ 1000.

- [input] integer param2

    Guaranteed constraints:
    -1000 ≤ param2 ≤ 1000.

- [output] integer
    The sum of the two inputs.
*/
func Add(param1 int, param2 int) int {
	switch {
	case param1 < -1000 || param2 < -1000 || param1 > 1000 || param2 > 1000:
		return 0
	default:
		return param1 + param2
	}
}

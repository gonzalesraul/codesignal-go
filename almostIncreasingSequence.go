package codesignal

/*
AlmostIncreasingSequence : Return true if it is possible to remove one element from the array in order to get a strictly increasing sequence, otherwise return false

Given a sequence of integers as an array, determine whether it is possible to obtain a strictly
increasing sequence by removing no more than one element from the array.

Example

    For sequence = [1, 3, 2, 1], the output should be
    almostIncreasingSequence(sequence) = false.

    There is no one element in this array that can be removed in order to get a strictly increasing sequence.

    For sequence = [1, 3, 2], the output should be
    almostIncreasingSequence(sequence) = true.

    You can remove 3 from the array to get the strictly increasing sequence [1, 2]. Alternately, you can remove 2 to get the strictly increasing sequence [1, 3].

Input/Output

- [execution time limit] 4 seconds (go)

- [input] array.integer sequence

    Guaranteed constraints:
    2 ≤ sequence.length ≤ 105,
    -105 ≤ sequence[i] ≤ 105.

- [output] boolean
    Return true if it is possible to remove one element from the array in order to get a strictly increasing sequence, otherwise return false.
*/
func AlmostIncreasingSequence(sequence []int) bool {
	count := len(sequence)
	if count < 2 || count > 100000 {
		return false
	}

	var gap int
	for index := 1; index < count; index++ {
		if sequence[index-1] >= sequence[index] {
			if gap++; gap > 1 {
				return false
			} else if index == 1 || index+1 == count {
				continue
			} else if sequence[index] > sequence[index-2] {
				sequence[index-1] = sequence[index-2]
			} else if sequence[index-1] >= sequence[index+1] {
				return false
			}
		}
	}

	return true
}

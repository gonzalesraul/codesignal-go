package codesignal

/*
CenturyFromYear : The number of the century the year is in.

Given a year, return the century it is in. The first century spans from the year 1 up to and including the year 100, the second - from the year 101 up to and including the year 200, etc.

Example

    For year = 1905, the output should be
    centuryFromYear(year) = 20;
    For year = 1700, the output should be
    centuryFromYear(year) = 17.

Input/Output

- [execution time limit] 4 seconds (go)

- [input] integer year

    A positive integer, designating the year.

    1 ≤ year ≤ 2005.
    Guaranteed constraints:

- [output] integer
    The number of the century the year is in.
*/
func CenturyFromYear(year int) (century int) /*Testing Named result */ {
	if year < 1 || year > 2005 {
		return 0
	}

	left := year % 100
	century = (year - left) / 100

	if left > 0 {
		century++
	}

	return
}

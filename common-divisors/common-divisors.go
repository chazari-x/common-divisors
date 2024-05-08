package common_divisors

import "slices"

func CommonDivisors(ints []int) []int {
	var divisors []int

	var m = slices.Min(ints)
	for i := 2; i <= m; i++ {
		divisors = append(divisors, i)
	}

	var commonDivisors []int

	for _, divisor := range divisors {
		if commonDivisorFinder(ints, divisor) {
			commonDivisors = append(commonDivisors, divisor)
		}
	}

	return commonDivisors
}

func commonDivisorFinder(ints []int, divisor int) bool {
	for _, i := range ints {
		if i%divisor != 0 {
			return false
		}
	}

	return true
}

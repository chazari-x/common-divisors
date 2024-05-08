package common_divisors

import "slices"

func CommonDivisors(ints []int) []int {
	// Поиск минимального числа в массиве
	var m = slices.Min(ints)
	// Поиск всех делителей числа m
	var divisors []int
	for i := 2; i <= m; i++ {
		divisors = append(divisors, i)
	}

	// Поиск общих делителей
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
		// Если число не делится на divisor, то это не общий делитель
		if i%divisor != 0 {
			return false
		}
	}

	return true
}

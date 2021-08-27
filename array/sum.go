package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numberToSum ...[]int) []int {
	lengthOfArgs := len(numberToSum)
	sums := make([]int, lengthOfArgs)

	for i, numbers := range numberToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTail(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}

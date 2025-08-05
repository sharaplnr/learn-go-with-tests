package main

func Sum(numbers[]int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lenghtResSlice := len(numbersToSum)
	// sums := make([]int, lenghtResSlice)
	var sums []int

	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}
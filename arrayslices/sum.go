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

func SumAllTails(numbersToSumTails ...[]int) []int {
	var sums []int
	
	for _, num := range numbersToSumTails {
		if len(num) > 1 {
			sums = append(sums, Sum(num[1:]))
		} else {
			sums = append(sums, 0)
		}
		
	}

	return sums
}
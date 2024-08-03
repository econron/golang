package arrsli

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tailIndex := len(numbers) - 1
		if tailIndex <= 0 {
			tailIndex = 0
		}
		tail := numbers[tailIndex:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
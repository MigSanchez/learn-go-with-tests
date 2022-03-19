package arrayslices

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSun ...[]int) (sums []int) {
	for _, numbers := range numbersToSun {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSun ...[]int) (sums []int) {
	for _, nums := range numbersToSun {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return
}

package arrays_and_slices

func Sum(arr []int) (sum int) {
	for _, n := range arr {
		sum += n
	}
	/* for i := 0; i < len(arr); i++ {
		sum += arr[i]
	} */
	return
}

func SumAll(numbers ...[]int) []int {
	// sums := make([]int, len(numbers))
	var sums []int

	for _, slice := range numbers {
		// sums[i] = Sum(slice)
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, slice := range numbers {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tails := slice[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
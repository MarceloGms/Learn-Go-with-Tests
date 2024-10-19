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
	sums := make([]int, len(numbers))

	for i, slice := range numbers {
		sums[i] = Sum(slice)
	}
	return sums
}
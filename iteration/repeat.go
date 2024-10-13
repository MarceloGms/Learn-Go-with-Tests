package iteration

func Repeat(x string, n int) (result string) {
	for i := 0; i < n; i++ {
		result += x
	}
	return
}
package iteration

func Repeat(x string) (result string) {
	for i := 0; i < 5; i++ {
		result += x
	}
	return
}
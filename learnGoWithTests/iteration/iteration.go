package iteration

func Repeat(input string, repeats int) (repeated string) {
	for i := 0; i < repeats; i++ {
		repeated += input
	}
	return repeated
}

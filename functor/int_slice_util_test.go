package functor

func intSlice(numInts int) []int {
	slice := make([]int, numInts)
	for i := 0; i < numInts; i++ {
		slice[i] = i + 5
	}
	return slice
}

func plusOne(i int) int {
	return i + 1
}
func minusOne(i int) int {
	return i - 1
}

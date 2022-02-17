package mathT

func SumArray(arr [10]int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
func SumSlice(s []int) int {
	var sum int
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

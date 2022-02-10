package mathT

func sumArray(arr [10]int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

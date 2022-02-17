package mathT

import "fmt"

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
func SumVariadic(vrd ...int) int {
	var sum int
	for i := 0; i < len(vrd); i++ {
		sum += vrd[i]
	}
	return sum
}

//for range
func SumSli_Range(s []int) int {
	var sum int
	for i, v := range s {
		fmt.Printf("index: %d | value: %d\n", i, v)
		sum += v //s[i] == v
	}
	return sum
}

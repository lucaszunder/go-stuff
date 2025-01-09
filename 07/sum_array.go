package main

import "fmt"

func main() {
	// [6,7,7,7]
	test := SumSlices([]int{1}, []int{1, 1}, []int{1, 1, 1}, []int{1, 1, 2})
	fmt.Println(test[1:])
}

func SumArray(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumSlices(slices_to_sum ...[]int) []int {
	size := len(slices_to_sum)
	output := make([]int, size)

	for i, v := range slices_to_sum {
		output[i] = SumArray(v)
	}

	return output
}

func SumSlices2(slices_to_sum ...[]int) []int {
	var output []int

	for _, v := range slices_to_sum {
		output = append(output, SumArray(v))
	}

	return output
}

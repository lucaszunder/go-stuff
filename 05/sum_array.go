package main

func SumArray(numbers [3]int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

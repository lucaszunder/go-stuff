package main

func SumArray(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

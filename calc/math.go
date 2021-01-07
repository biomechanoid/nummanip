package calc

// Add func
func Add(numbers ...int) int {
	var sum int
	for _, i := range numbers {
		sum = sum + i
	}

	return sum
}

package calc

import "errors"

// Add func
func Add(numbers ...int) (int, error) {
	var sum int
	if len(numbers) < 2 {
		return 0, errors.New("Provide at least 2 numbers")
	} else {
		for _, i := range numbers {
			sum = sum + i
		}
	}

	return sum, nil
}

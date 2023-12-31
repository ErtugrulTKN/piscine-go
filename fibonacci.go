package piscine

func Fibonacci(value int) int {
	if value < 0 {
		return -1
	}

	if value == 0 {
		return 0
	} else if value == 1 {
		return 1
	}

	return Fibonacci(value-1) + Fibonacci(value-2)
}

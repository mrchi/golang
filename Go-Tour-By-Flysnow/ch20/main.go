package main

var cache = make(map[int]int, 20)

func Fibonacci(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	result := 0

	switch {
	case n < 0:
		return 0
	case n == 0:
		return 0
	case n == 1:
		return 1
	default:
		result = Fibonacci(n-1) + Fibonacci(n-2)
	}
	cache[n] = result
	return result
}

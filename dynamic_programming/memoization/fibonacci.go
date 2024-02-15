package memoization

var fibCache = make(map[int]int)

func Fibonacci(n int) int {


	if result, found := fibCache[n]; found {
		return result
	}

	if n <= 1 {
		return n
	}

	fibValue := Fibonacci(n-1) + Fibonacci(n-2)

	fibCache[n] = fibValue

	return fibValue
}

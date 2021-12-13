package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	result := 1.
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}

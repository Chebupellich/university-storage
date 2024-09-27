package mathutils

func Factorial(n int) int {
	acc := 1
	for i := 1; i <= n; i++ {
		acc *= i
	}

	return acc
}

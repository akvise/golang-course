package fibonacci

func FibLoop(n int) int {
	a, b :=  1, 0
	for i := 0; i < n - 1; i++ {
		b, a = a, a + b
	}
	return a
}

func FibRecursion(n int) int {
	if n == 1 || n == 2{
		return 1
	}
	return FibRecursion(n-2) + FibRecursion(n-1)
}


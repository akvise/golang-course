package main

func average(a[]float32) float32 {
	var sum float32
	for i := range a{
		sum += a[i]
	}
	return sum / float32(len(a))
}

func main() {
	a := []float32{1,2,3,4,5,8}
	average(a)
}

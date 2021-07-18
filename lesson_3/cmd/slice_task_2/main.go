package main

import "fmt"

func reverse(a[]int64) []int64 {
	b := make([]int64, len(a), cap(a))
	for i := range a {
		b[i] = a[len(a)-1 - i] 		//reverse index in "a" slice
	}
	return b
}

func main() {
	a := []int64{1,2,5,15}
	fmt.Println(reverse(a))
}

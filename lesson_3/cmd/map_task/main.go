package main

import (
	"fmt"
	"sort"
)

func PrintSorted(a map[int]string) {
	keys := make([]int, 0, len(a))
	for i := range a {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for i := range keys {
		fmt.Print(a[i], " ")
	}
}

func main() {
	a := map[int]string{1: "b", 0: "a", 2: "c"}
	PrintSorted(a)
}

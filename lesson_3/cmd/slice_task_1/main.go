package main

import (
	"fmt"
	"unicode/utf8"
)
func max(a[]string) string {
	var maxlength, index int
	for i := range a{
		if utf8.RuneCountInString(a[i]) > maxlength {
			maxlength = utf8.RuneCountInString(a[i])
			index = i
		}
	}
	return a[index]
}

func main() {
	a := []string{"one", "two", "three"}
	fmt.Println(a)
	fmt.Println("The longest word in slice:", max(a))
}

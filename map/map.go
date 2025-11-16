package main

import "fmt"

func calculate(a []int) map[int]int {
	number := make(map[int]int)
	for i, value := range a {
		number[i] = value
	}
	return number
}
func main() {
	a := []int{1, 2, 2, 3, 3}
	b := calculate(a)
	fmt.Println(a)
	fmt.Println(b)
}

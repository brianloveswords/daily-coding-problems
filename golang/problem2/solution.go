package main

import "fmt"

// Solution is the solution to the problem
func Solution(xs []int) []int {
	out := make([]int, len(xs))
	total := 1
	for _, val := range xs {
		total *= val
	}
	for i, val := range xs {
		out[i] = total / val
	}
	fmt.Println("{}", out)
	return out
}

// NoDivisionSolution is the solution to the problem but doesn't use division
func NoDivisionSolution(xs []int) []int {
	out := make([]int, len(xs))
	for i := range xs {
		out[i] = 1
		for j, val := range xs {
			if j == i {
				continue
			}
			out[i] *= val
		}
	}
	fmt.Println("{}", out)
	return out
}

func main() {
	out := Solution([]int{4, 5, 6})
	fmt.Println("{}", out)
}

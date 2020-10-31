package main

import (
	"fmt"
	"log"
	"os"
)

// PerfectNumber returns the nth perfect number. A number is "perfect" if its
// digits sum to 10. (Note: this is not really a perfect number)
func PerfectNumber(n int) int {
	logger := log.New(os.Stderr, "debug: ", log.LstdFlags|log.Lmicroseconds)
	if n <= 0 || n > 59271 {
		panic(fmt.Sprintf("valid range for n is [1, 59271] got %d", n))
	}
	i := 0
	current := 19
	nextPower := 100
	for {
		if isPerfectFast(current) {
			i++
		}
		if i == n {
			return current
		}
		if current == nextPower {
			logger.Printf("passed %d\n", nextPower)
			nextPower *= 10
		}
		current++
	}
}

/*
	`isPerfect` and frieds was the 1st (really 2nd*) attempt at a naive
	solution. It is slow: it takes several minutes to find the terminal
	perfect number.

	*: the first, _very_ naïve pass converted the number to a string, then
	split it, then converted each part back to an integer to sum them up. I
	knew this was very wasteful, but it helped me sketch out the general shape
	of the problem.

*/
func isPerfect(n int) bool {
	parts := intToParts(n)
	return sumParts(parts) == 10
}

func sumParts(xs []int) int {
	var sum int = 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

func intToParts(n int) []int {
	if n < 10 {
		return []int{n}
	}
	parts := make([]int, 0)
	for {
		remainingPlaces := n / 10
		if remainingPlaces == 0 {
			parts = append(parts, n)
			return parts
		}
		parts = append(parts, n-remainingPlaces*10)
		n = remainingPlaces
	}
}

/*
	`isPerfectFast` does the same thing as isPerfect, only it smashes
	everything together in place rather than allocating arrays and passing
	those around.
*/
func isPerfectFast(n int) bool {
	sum := 0
	for {
		remainingPlaces := n / 10
		if remainingPlaces == 0 {
			sum += n
			return sum == 10
		}
		sum += n - remainingPlaces*10
		n = remainingPlaces
	}
}

func main() {
	fmt.Println(PerfectNumber(59271))
}

package main

import (
	"fmt"  // for reading from stdin and printing to stdout
)

func main() {
	// Read two integers, n and m, from stdin
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	// Initialize result to 0
	result := 0

	// Loop over the input `n` times
	for i := 0; i < n; i++ {
		// Read `m` integers from stdin
		data := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &data[j])
		}

		// Find the minimum value in the slice
		minValue := data[0]
		for _, num := range data {
			if num < minValue {
				minValue = num
			}
		}

		// Update the result if the current minimum value is larger
		if minValue > result {
			result = minValue
		}
	}

	// Print the final result
	fmt.Println(result)
}

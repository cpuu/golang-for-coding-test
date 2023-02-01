package main

import (
	"fmt"
)

func main() {
	var n, k int
	// Read `n` and `k` from the input, separated by a space
	fmt.Scanf("%d %d", &n, &k)

	result := 0
	for {
		// Calculate the target by dividing `n` by `k` and multiplying the result by `k`
		target := (n / k) * k
		// Add the difference between `n` and the target to the `result`
		result += (n - target)
		// Set `n` to the target
		n = target
		// Break the loop if `n` is less than `k`
		if n < k {
			break
		}
		// Increment `result` by 1
		result++
		// Divide `n` by `k`
		n /= k
	}
	// Add the difference between `n` and 1 to the `result`
	result += (n - 1)
	// Print the final value of `result`
	fmt.Println(result)
}

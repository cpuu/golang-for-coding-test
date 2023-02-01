package main

import (
	"fmt"
	"sort"
)

func main() {
	// Read N, M, K from stdin, separated by spaces
	var n, m, k int
	fmt.Scanf("%d%d%d", &n, &m, &k)

	// Read N numbers from stdin, separated by spaces
    v := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &v[i])
    }

	// Sort the input numbers
	sort.Ints(v)
	first := v[n-1] // Get the largest number
	second := v[n-2] // Get the second largest number

	// Calculate the number of times the largest number will be added
	cnt := (m / (k + 1)) * k
	cnt += m % (k + 1)

	// Calculate the final answer
	result := 0
	result += cnt * first
	result += (m - cnt) * second

	// Print the final answer
	fmt.Println(result)
}
  
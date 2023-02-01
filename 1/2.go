package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize random number generator

	// Insert 10,000 integers into the array
	var array []int
	for i := 0; i < 10000; i++ {
		array = append(array, rand.Intn(100)+1) // Random integer between 1 and 100
	}

	// Measure performance of selection sort program
	startTime := time.Now()

	// Selection sort program source code
	for i := 0; i < len(array); i++ {
		minIndex := i // Index of smallest element
		for j := i + 1; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i] // Swap
	}

	endTime := time.Now() // End measurement
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Selection sort performance measurement:", elapsedTime) // Print execution time

	// Reinitialize the array with random data
	array = nil
	for i := 0; i < 10000; i++ {
		array = append(array, rand.Intn(100)+1) // Random integer between 1 and 100
	}

	// Measure performance of the default sort library
	startTime = time.Now()

	// Use the default sort library
	sort.Ints(array)

	endTime = time.Now() // End measurement
	elapsedTime = endTime.Sub(startTime)
	fmt.Println("Default sort library performance measurement:", elapsedTime) // Print execution time
}

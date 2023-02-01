package main

import "time"

func main() {
    startTime := time.Now() // Start measurement

    // Program source code
    endTime := time.Now() // End measurement
    elapsedTime := endTime.Sub(startTime)
    println("time:", elapsedTime) // Print execution time
}

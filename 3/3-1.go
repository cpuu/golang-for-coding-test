package main

import "fmt"

func main() {
    n := 1260
    count := 0
    coinTypes := []int{500, 100, 50, 10}

    for _, coin := range coinTypes {
        count += n / coin
        n %= coin
    }

    fmt.Println(count)
}
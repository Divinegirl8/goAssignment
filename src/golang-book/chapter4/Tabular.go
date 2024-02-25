package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("N   N2   N3   N4")

    for count := 1; count <= 5; count++ {
        number := math.Pow(float64(count), float64(count))
        fmt.Printf("%-5d%-5d%-5d%-5d\n", count, int(number), int(math.Pow(float64(count), 3)), int(math.Pow(float64(count), 4)))
    }
}


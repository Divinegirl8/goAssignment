package main

import (
    "fmt"
    "math"
)

func pythagorean(side1, side2, side3 int) {
    for count1 := 1; count1 <= side1; count1++ {
        for count2 := 1; count2 <= side2; count2++ {
            for count3 := 1; count3 <= side3; count3++ {
                if math.Pow(float64(count1), 2)+math.Pow(float64(count2), 2) == math.Pow(float64(count3), 2) {
                    fmt.Printf("%d^2 + %d^2 = %d^2\n", count1, count2, count3)
                }
            }
        }
    }
}

func main() {
    pythagorean(20, 20, 20)
}


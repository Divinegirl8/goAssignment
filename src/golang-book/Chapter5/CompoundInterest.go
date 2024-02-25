package main

import (
    "fmt"
    "math"
)

func main() {
    var principal int = 1000
    var penny int = principal * 100
    var rate float64 = 5.0 / 100 
    var compoundingFrequently int = 4
    annual := 0
    for count := 1; count <= 10; count++ {
        annual = int(float64(penny) * math.Pow(1+rate/float64(compoundingFrequently), float64(compoundingFrequently*count)))
    }

    var dollar = annual / 100
    var cent = annual % 100

    fmt.Printf("The compounded interest is $%d.%d cents\n", dollar, cent)
}


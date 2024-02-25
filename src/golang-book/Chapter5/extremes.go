package main

import (
    "fmt"
)

func main() {
    sum := 0
    var max, min int

    max = -2147483648
  
    min = 2147483647

    for count := 0; count < 10; count++ {
        fmt.Println("Enter a number: ")
        var number int
        fmt.Scanf("%d", &number)

        if number > max {
            max = number
        }
        if number < min {
            min = number
        }
    }

    sum = max + min
    fmt.Println("The sum of the extremes is", sum)
}


package main

import (
    "fmt"
)

func question() []int {
    var numbers [5]int

    for index := 0; index < len(numbers); index++ {
        var input int
        fmt.Printf("Enter number %d (between 1 - 30): ", index+1)
        fmt.Scanln(&input)
       
        for input < 1 || input > 30 {
            fmt.Println("Invalid input. Please enter a number between 1 - 30.")
            fmt.Printf("Enter number %d (between 1 - 30): ", index+1)
            fmt.Scanln(&input)
        }
        numbers[index] = input
    }

    return numbers[:]
}

func main() {
    result := question()
   
   for index := 0; index < len(result); index++{
       for count := 0; count < result[index]; count++{
       fmt.Print("*")
       }
       fmt.Println()
   }
}


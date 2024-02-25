


package main

import (
"fmt"
"strings"
"strconv"
)

type Palindrome struct {}

func (palindrome Palindrome) isPalindrome(number int) bool{
convertToString := strconv.Itoa(number)
var stringBuilder strings.Builder


for index := len(convertToString)-1; index >= 0; index--{
stringBuilder.WriteByte(convertToString[index])
}
reversedString := stringBuilder.String()

return reversedString == convertToString
}

func question() int{
fmt.Println("Enter a number that consist of four length: (e.g 1222): ")
var input int
fmt.Scanf("%d",&input)

return input
}


func main() {
palindrome := Palindrome{}

input := question()


   for input < 1000 || input > 9999 {
        fmt.Println("Please enter a four-digit number.")
        input = question()
    }


result := palindrome.isPalindrome(input)

if result {
fmt.Println("The number you entered is a plandrome")
}else{
fmt.Println("The number you entered is not a plandrome")
}




}

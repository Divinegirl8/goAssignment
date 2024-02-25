package main

import "fmt"

func main(){

fmt.Println("Enter a number: ")
var firstNumber int
fmt.Scanf("%d",&firstNumber)


fmt.Println("Enter another number: ")
var secondNumber int
fmt.Scanf("%d",&secondNumber)


if firstNumber == secondNumber {
fmt.Println(0)
}else if firstNumber > secondNumber {
fmt.Println(1)
}else if firstNumber < secondNumber {
fmt.Println(-1)
}


}

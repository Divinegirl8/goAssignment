package main

import "fmt"

func main(){
largest := 0
secondLargest := 0

for count:=0; count < 10; count++{
fmt.Println("Enter a number: ")
var input int
fmt.Scanf("%d",&input)

if largest < input{
secondLargest = largest
largest = input
}else if secondLargest < input{
secondLargest = input
}

}
fmt.Printf("The largest is %d\nThe second largest is %d\n",largest,secondLargest)
}

package main

import "fmt"

type LargestNumber struct {
max int
}


func (ln *LargestNumber) findLargestNumber(number int)int{


if (number > ln.max){
ln.max = number
}

return ln.max;
} 


func main(){

large := LargestNumber{}
count := 0

for count < 10{
fmt.Println("Enter number: ")
var number int
fmt.Scanf("%d",&number)


large.findLargestNumber(number)

count += 1
 
}

fmt.Printf("The largest number is %d\n",large.max)
}

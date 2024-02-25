package main

import "fmt"

type RightAngle struct {}

func (angle RightAngle) question() int{
fmt.Println("enter the length of the base of a triangle: (between 1 - 10)")
var length int
fmt.Scanf("%d",&length)

return length

}

func main(){

aesterik := ""


triangle := RightAngle{}

length := triangle.question()

for length < 1 || length > 10{
fmt.Println("Error!!!: " )
length = triangle.question()

}


for count := 0; count < length; count++{
aesterik += "*"
fmt.Println(aesterik);




}}

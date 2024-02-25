package main

import "fmt"


func main(){
fmt.Println("Enter a number")
var num int
fmt.Scanf("%d",&num)


sum := 0

for{
fmt.Println("Enter a number")
var secondNum int
fmt.Scanf("%d",&secondNum)

sum += secondNum


if sum == num || sum > num{
break
}

}



}

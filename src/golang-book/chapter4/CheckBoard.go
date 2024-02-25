package main

import "fmt"

func main(){

for index := 0; index < 5; index++{
for row1 := 0; row1 < 8; row1++{
fmt.Print("*" + " ")
}

fmt.Println()
fmt.Print(" ")

for row2 := 0; row2 < 8; row2++{
fmt.Print("*" + " ")
}
fmt.Println()

}}

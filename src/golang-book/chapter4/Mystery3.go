package main

import "fmt"

func main(){

row := 5

for row >= 1{
column := 5

for column >= 1 {

if column % 2 ==0{
fmt.Print("X")
} else{
fmt.Print("O")
}

column -= 1 // The initial code is column += 1 and i changed it to column -= 1 because the initial one will result to an infinite loop 
}
row -= 1
fmt.Println()
}
}

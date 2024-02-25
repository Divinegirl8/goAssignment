package main

import "fmt"


func main(){
x := 2

x += x++ - 5;

fmt.Println(x) // error

}

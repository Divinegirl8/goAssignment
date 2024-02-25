package main

import "fmt"

type TaxCalculator struct{}

func (tc TaxCalculator) calculateTax(percent int,amount int)int{
return (percent * amount) / 100
}

func main(){

for count := 1; count <= 3;  count++{

fmt.Println("Enter citizen's name: ")
var name string
fmt.Scanf("%s",&name)

fmt.Println("Enter earnings: ")
var earning int
fmt.Scanf("%d",&earning)

var totalTax int
cal := TaxCalculator{}

if (earning <= 30_000){

totalTax = cal.calculateTax(15,earning)
fmt.Printf("Total tax is %d\n", totalTax)
}else{
totalTax = cal.calculateTax(20,earning)
fmt.Printf("Total tax is %d\n", totalTax)
}



}


}



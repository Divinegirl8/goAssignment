package main

import "fmt"

type Gas struct {}

func (gas Gas) calcAverage(totalGallon float64, totalMiles float64) float64{
return totalGallon / totalMiles

}

func main(){

var gallonTotal float64
var milesTotal float64




for{

fmt.Println("Enter number of Gallons used for tankful: ")
var gallons float64
fmt.Scanf("%f", &gallons)

if gallons != -1{
gallonTotal += gallons}

if gallons == -1{
break
}

fmt.Println("Enter number of Miles driven: ")
var miles float64
fmt.Scanf("%f",&miles)
milesTotal += miles

sum := gallons / miles

fmt.Println("The cost of gallons used per miles is ",sum)


}
var total float64
gas := Gas{}
 
total = gas.calcAverage(gallonTotal,milesTotal)

fmt.Println("The average of MilesPerGallon is ",total)

}

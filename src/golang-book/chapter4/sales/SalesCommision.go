package main

import (
"bufio"
"fmt"
"os"
 "strings"
)


type SalesCommission struct{}


func (sales SalesCommission) calculateGrossSales(percent float64,sum float64) float64{

return (percent * sum) / 100

}


func (sales SalesCommission) calculateWeeklySalary(salary float64, percent float64,sum float64) float64{
sc := SalesCommission{}

var grossSales float64

grossSales = sc.calculateGrossSales(percent,sum)

return grossSales + salary

}


func (sales SalesCommission) userInput(){
var sum float64
salary := 200.0
sc := SalesCommission{}

for{
fmt.Println("Enter amount of item value for a salesperson (press -2 to exit): ")
var amount float64
fmt.Scanf("%f\n",&amount)

if (amount != -2){
sum += amount
}

if (amount == -2){
break
}

}

var weeklySalary float64

weeklySalary = sc.calculateWeeklySalary(salary, 9, float64(sum))
fmt.Printf("Sales person receives: $%.2f\n",weeklySalary)

}


func (sales SalesCommission) question() string{
  scanner := bufio.NewScanner(os.Stdin)
 fmt.Println("Do you want to continue? ")
 scanner.Scan()
 
 return strings.ToLower(scanner.Text())
}


func main() {
    sc := SalesCommission{}

    sc.userInput()
   
    for {
        choice := sc.question()
        if choice == "yes" {
            sc.userInput()
        } else if choice != "no" {
            fmt.Println("wrong input choose between yes or no")
        } else if choice == "no" {
            break
        }
    }
}


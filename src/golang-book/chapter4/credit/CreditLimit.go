package main

import "fmt"

type CalCredit struct{}

func (c CalCredit) calculateCredit(bal int, charges int, credit int)int{
return  (bal + charges) - credit
}

func main(){
fmt.Println("Enter account number: ")
var acctNumber int
fmt.Scanf("%d", &acctNumber)

fmt.Println("Enter balance at the beginning of the month: ")
var bal int
fmt.Scanf("%d", &bal)


fmt.Println("total of all items charged by the customer this month: ")
var charges int
fmt.Scanf("%d", &charges)

fmt.Println("total of all credits applied to the customer’s account this month")
var credit int
fmt.Scanf("%d", &credit)

fmt.Println("allowed credit limit.")
var limit int
fmt.Scanf("%d",&limit)


cal := CalCredit{}

newBalance := cal.calculateCredit(bal,charges,credit)

if newBalance > limit{
fmt.Printf("NewBalance is %d\nCredit limit exceeded", newBalance)
} else{
fmt.Printf("NewBalance is %d\n", newBalance)
}

}

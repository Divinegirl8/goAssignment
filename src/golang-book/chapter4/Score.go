package main

import "fmt"


func main(){
pass := 0
fail := 0
count := 0


for count < 10{
fmt.Println("Enter student score (pass = 1 and fail = 2): ")
var score int
fmt.Scanf("%d",&score)

if score == 1{
pass += 1
count += 1
}

if score == 2{
fail += 1
count += 1
}

}

fmt.Printf("The number of student that passed is %d\nThe number of student that failed is %d\n",pass,fail)

if(pass > 8){
fmt.Println("Bonus to the instructor")
}


}

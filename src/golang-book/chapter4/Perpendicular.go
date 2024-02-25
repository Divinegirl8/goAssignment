package main


import "fmt"


func main(){

fmt.Println("Enter first x axis: ")
var firstX int
fmt.Scanf("%d",&firstX)


fmt.Println("Enter first y axis: ")
var firstY int
fmt.Scanf("%d",&firstY)


fmt.Println("Enter second x axis: ")
var secondX int
fmt.Scanf("%d",&secondX)


fmt.Println("Enter second y axis: ")
var secondY int
fmt.Scanf("%d",&secondY)


  if (firstX == secondX) {
            fmt.Println(firstX , "and" , secondX , "are located on a line perpendicular to an axis")
  } else {
            fmt.Println(firstX, "and" , secondX , "are not located on a line perpendicular to an axis")
        }

 if (firstY == secondY ) {
            fmt.Println(firstY , "and" , secondY , "are located on a line perpendicular to an axis")
        }else {
            fmt.Println(firstY , "and" , secondY , "are not located on a line perpendicular to an axis")
        }


}

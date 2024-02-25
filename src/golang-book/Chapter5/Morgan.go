package main

import "fmt"


func main(){

i := 2;
j := 3;
k := 2;
m := 2;

 fmt.Println(i == 2);
 fmt.Println(j == 5);
 fmt.Println((i >= 0) && (j <= 3));
 fmt.Println((m <= 100) && (k <= m));
 fmt.Println((j >= i) || (k != m));
 fmt.Println((k + i < j) || (4 - j >= k));
 fmt.Println(!(k > j));
}

package main

import "testing"

func TestLargest(t *testing.T){
largest := LargestNumber{max:0}

TestCase := [] struct{
number int
expected int
}{
{10,10},

}


for _,tc := range TestCase{
result := largest.findLargestNumber(tc.number)

if result != tc.expected{
t.Errorf("Actual : %d\nExpected : %d\n",result,tc.expected)
}
}


}

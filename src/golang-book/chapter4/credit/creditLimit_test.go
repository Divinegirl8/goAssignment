package main

import "testing"

func TestCreditLImit(t *testing.T){
creditLimit :=  CalCredit{}

TestCase := []  struct{
bal     int
charges int
credit  int
expected int
}{
{11,11,11,11},
{-34,-56,-10,-80},
}

for _,tc := range TestCase{
result := creditLimit.calculateCredit(tc.bal,tc.charges,tc.credit)
if result != tc.expected{
t.Errorf("Actual : %d\nExpected : %d",result,tc.expected)
}
}

}



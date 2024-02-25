package main

import "testing"


func TestSales(t *testing.T){
sc := SalesCommission{}

testCases := [] struct{
salary float64
percent float64
sum float64
expected float64
}{{200,9,5000,650},}

for _,tc := range testCases{
result := sc.calculateWeeklySalary(tc.salary,tc.percent,tc.sum)

if result != tc.expected{
t.Errorf("Actual : %f\nExpected : %f\n",result,tc.expected)
}

}



}

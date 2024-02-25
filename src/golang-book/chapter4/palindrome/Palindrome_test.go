package main

import "testing"


func TestPalindromet(t *testing.T){
palindrome := Palindrome{}

testCases := [] struct {
number int
expected bool
}{
{1991,true},
{2000, false},
}


for _, tc := range testCases{

result := palindrome.isPalindrome(tc.number)
if result != tc.expected{
t.Errorf("Actual : %t\nExpected : %t\n ",result,tc.expected)
}
}
}

package main

import (
    "testing"
)

func TestCalc(t *testing.T) {
    calculator := TaxCalculator{}

    testCases := []struct {
        percent  int
        amount   int
        expected int
    }{
        {15, 30000, 4500},
    }

    for _, tc := range testCases {
        result := calculator.calculateTax(tc.percent, tc.amount)
        if result != tc.expected {
            t.Errorf("Actual: %d\nExpected: %d", result, tc.expected)
        }
    }
}
// go mod init <module name>


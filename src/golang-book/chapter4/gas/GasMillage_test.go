package main

import ("testing")

func TestGas(t *testing.T){
gas := Gas{}

testCases := [] struct{
totalGallon float64
totalMiles float64
expected float64
}{{10,2,5},
{9,3,3},
}


for _, tc := range testCases{
result := gas.calcAverage(tc.totalGallon,tc.totalMiles)

if result != tc.expected {
t.Errorf("Actual : %f\nExpected : %f\n",result,tc.expected)
}
}

}

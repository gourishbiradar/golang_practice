package main

import (
	"fmt"
	"testing"
	. "fractions/fraction_multiplication"
)

//Testing
func TestComputeGCD(t *testing.T) {
	got := ComputeGCD(6,10)
	if got != 2 {
		t.Errorf("ComputeGCD(6,10) = %d ; want 2", got)
	}
}

func main() {
	a := Fraction{Numerator: 6, Denominator: 10}
	a.Reduce()
	fmt.Println(a)

	b := Fraction{Numerator: 4,Denominator: 10}
	a.Multiply(b)
	fmt.Println(a,b)
}

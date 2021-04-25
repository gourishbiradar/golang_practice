package fractions


// To be treated as a class, to represent fractions
type Fraction struct {
	Numerator, Denominator int
}

func (f *Fraction) Reduce() {
	gcd := ComputeGCD(f.Numerator,f.Denominator)
	f.Numerator /= gcd
	f.Denominator /= gcd
}

func (a *Fraction) Multiply(b Fraction) {
	a.Numerator = a.Numerator * b.Numerator
	a.Denominator = a.Denominator * b.Denominator
	a.Reduce()
}

// Utility Function: Recursively compute GCD
func ComputeGCD(Numerator int, Denominator int) (gcd int) {
	if((Numerator % Denominator) == 0){
		return Denominator
	}
	return ComputeGCD(Denominator,Numerator%Denominator)
}





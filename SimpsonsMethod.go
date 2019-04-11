package main

import (
"math"

"golang.org/x/exp/errors/fmt"
)

const EPS = 0.00001
const A = 0.
const d = -3.
const c = 2.

func F(x float64) float64 {
	return math.Pow(math.E, d*x) * math.Cos(c*x)
}

func F2(x float64) float64 {
	return math.Pow(math.E, d*x) * math.Sin(c*x)
}

func points(a, b float64, n int) []float64 {
	x := make([]float64, 2*n+1)
	h := (b - a) / (2 * float64(n))
	for i := 0; i <= 2*n; i++ {
		x[i] = a + float64(i)*h
	}
	return x
}

func quadraticFormula(x1, x2, x3 float64) float64 {
	return ((x3 - x1) / 6) * (F(x1) + 4*F(x2) + F(x3))
}

func quadraticFormula2(x1, x2, x3 float64) float64 {
	return ((x3 - x1) / 6) * (F2(x1) + 4*F2(x2) + F2(x3))
}

func SimpsonsMethod(a, b float64, n int) float64 {
	X := points(a, b, n)
	result := 0.
	for i := 0; i < 2*n; i = i + 2 {
		result += quadraticFormula(X[i], X[i+1], X[i+2])
	}
	return result
}

func SimpsonsMethod2(a, b float64, n int) float64 {
	X := points(a, b, n)
	result := 0.
	for i := 0; i < 2*n; i = i + 2 {
		result += quadraticFormula2(X[i], X[i+1], X[i+2])
	}
	return result
}

func findUpperBound(n int) float64 {
	H := 10000.
	B := A + 1
	for math.Abs(SimpsonsMethod(B, H, n)) >= EPS/2 {
		B++
	}
	return B
}

func methodRunge(a, b float64, m, n int) int {
	Ih := SimpsonsMethod(a, b, n)
	Ih2 := SimpsonsMethod(a, b, 2*n)
	R0h2 := (Ih2 - Ih) / (math.Pow(2, float64(m)) - 1)
	//R0h2 := (Ih2 - Ih) / 15
	if math.Abs(R0h2) > EPS {
		return methodRunge(a, b, m, 2*n)

	}
	return n
}

func ApriorMark(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	Integral1 := SimpsonsMethod(a, b, n)
	Integral2 := SimpsonsMethod2(a, b, n)
	return (math.Pow(h, 4) / 2880) * (-109*Integral1 + 120*Integral2)
}

func pointsWithApriorMark(a, b float64, n int) int {
	AprM := ApriorMark(a, b, n)
	if math.Abs(AprM) > EPS {
		return pointsWithApriorMark(a, b, 2*n)
	}
	return n
}

func m(a, b float64, n int) float64 {
	Ih := SimpsonsMethod(a, b, n)
	Ih2 := SimpsonsMethod(a, b, 2*n)
	Ih4 := SimpsonsMethod(a, b, 4*n)
	m := (Ih2 - Ih) / (Ih4 - Ih2)
	return math.Log2(m)
}

func main() {
	I := 3. / 13
	fmt.Printf("I= : %.10f ", I)
	fmt.Print("\n")

	B := findUpperBound(2000)
	fmt.Printf("Upper bound: %v ", B)
	fmt.Print("\n")

	m := int(m(A, B, 1))
	N := methodRunge(0, B, m, 1)

	fmt.Print(SimpsonsMethod(A, B, N))
	fmt.Print("\n")

	Rh := I - SimpsonsMethod(A, B, N)
	fmt.Printf("Slingshot= %0.10f", Rh)
	fmt.Print("\n")

	fmt.Printf("h= %0.10f", (B-A)/float64(N))
	fmt.Print("\n")

	N2 := pointsWithApriorMark(A, B, 1)
	fmt.Printf("ApriorMark points= %v", N2)
	fmt.Print("\n")

	Rh2 := I - SimpsonsMethod(A, B, N2)
	fmt.Printf("Slingshot= %0.10f", Rh2)
	fmt.Print("\n")
}

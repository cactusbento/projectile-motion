package main
import "math"


func correctTheta(v float64, x float64, y float64) (float64, float64) { 
	const g float64 = 9.82


	v2 := math.Pow(v, 2)
	v4 := math.Pow(v, 4)
	x2 := math.Pow(x, 2)

	usqrt := v4 - g*( g*x2 + 2*y*v2)
	numA := v2 + math.Sqrt(usqrt)
	numB := v2 - math.Sqrt(usqrt)
	
	return math.Atan2(numB , (g*x)) * (180/math.Pi), math.Atan2(numA , (g*x)) * (180/math.Pi)
}

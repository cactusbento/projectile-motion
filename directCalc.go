package main
import "math"


func correctTheta(v float64, x float64, y float64) (float64, float64) { 
	const g float64 = 9.82


	v2 := v * v
	v4 := math.Pow(v, 4)
	x2 := x * x 

	usqrt := v4 - g*( g*x2 + 2*y*v2)
	numA := v2 + math.Sqrt(usqrt)
	numB := v2 - math.Sqrt(usqrt)
	
	return math.Atan2(numB , (g*x)) * (180/math.Pi), math.Atan2(numA , (g*x)) * (180/math.Pi)
}

func maxHeight(v float64, theta float64) float64 {
	const g float64 = 9.82 
	rad := theta * (math.Pi / 180)

	v2 := v * v 
	sT := math.Sin(rad) * math.Sin(rad)

	return (v2 * sT) / (2 * g)
}


func timeOfFlight(v float64, a float64, yin float64) float64 {
	const g float64 = 9.82
	rad := a * (math.Pi/180)
	y 	:= -yin


	flight := v * math.Sin(rad)
	flight2 := flight * flight
	num := flight + math.Sqrt( flight2 + 2*g*y )
	

	return (num / g)
}

func vDispToX(v, x, theta float64) float64 { 
	const g float64 = 9.82
	rad := theta * (math.Pi / 180)

	v2 := v * v 
	x2 := x * x 
	cos2 := math.Cos(rad) * math.Cos(rad)

	p1 := math.Tan(rad) * x 
	p2 := 2 * v2 * cos2 

	return p1 - (g / p2) * x2 
}

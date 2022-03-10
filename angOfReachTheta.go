package main 

import "math"

func angOfReach(v float64, dist float64) (float64, float64) {
	var d float64 = dist
	const g float64 = 9.82

	paren := (g*d) / math.Pow(v,2)

	shallow := 0.5*math.Asin(paren)
	steep := 0.5*math.Acos(paren)

	return shallow * (180/math.Pi), steep * (180/math.Pi) + 45
}

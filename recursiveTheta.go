package main

import (
	"math"
)

func recursiveTheta(v float64, loft bool, min float64, max float64, x float64, y float64) float64 {
	const g float64 = 9.82
	const ERROR float64 = 0.000000001


	if min > max {
		return -1
	}

	mid := (min+max) / 2
	midRad := mid * (math.Pi/180)
	

	finalD := ((v*v)/g)*math.Sin(2*midRad)
	finalY := math.Tan(midRad) * x - g / (2*v*v*math.Pow(math.Cos(midRad),2) ) * x*x

	if min == max {return -1}
	if !loft {
		if finalY > y + ERROR {
			return recursiveTheta(v, loft, min, mid, x, y)
		} else if finalY < y - ERROR {
			return recursiveTheta(v, loft, mid, max, x, y)
		}
	} else {
		if finalD < x - ERROR {
			return recursiveTheta(v, loft, min, mid, x, y)
		} else if finalD > x + ERROR {
			return recursiveTheta(v, loft, mid, max, x, y)
		}
	}
	return mid
}


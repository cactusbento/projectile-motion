package main

import (
	"math"
)


func timeOfFlight(v float64, a float64, yin float64) float64 {
	const g float64 = 9.82
	var rad float64 = a * (math.Pi/180)
	var y float64 = -yin


	var flight float64 = v * math.Sin(rad)
	var num float64 = flight + math.Sqrt( math.Pow(flight,2) + 2*g*y )
	

	return (num / g)
}

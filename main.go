package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type coord struct {
	X	float64
	Y	float64
}

func (c coord) print(name string) {
	fmt.Print(name, ": (" , c.X , ", " , c.Y , ")\n" )
}

func main() {
	if len(os.Args) < 4 {
		panic("3 arguments required: Velocity, Distance, Height")
	}

	var numbers []float64

	for _,arg := range os.Args[1:] {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			numbers = append(numbers, n)
		}
	}


	relative := coord{ numbers[1], numbers[2] }
	var INIT_VEL float64 = numbers[0]

	var distance float64 = math.Sqrt( math.Pow(relative.X,2) + math.Pow(relative.Y,2) )

	cA, cB := correctTheta(INIT_VEL, relative.X, relative.Y)

	aorA, aorB := angOfReach(INIT_VEL, distance) 
	aorA += math.Atan(relative.Y / relative.X) * (180/math.Pi)
	aorB += math.Atan(relative.Y / relative.X) * (180/math.Pi)
	
	rA, rB := recursiveTheta(INIT_VEL, false, 0, 45, relative.X, relative.Y),
				recursiveTheta(INIT_VEL, true, 45, 90, relative.X, relative.Y)

	fmt.Println("Correct Shallow and Steep angles:",cA,",",cB)
	fmt.Println("Angle of reach Shallow, Steep :",aorA,",",aorB)
	fmt.Println("Recursive Shallow, Steep :",rA,",",rB)

	cTA, cTB := timeOfFlight(INIT_VEL, cA, relative.Y), timeOfFlight(INIT_VEL, cB, relative.Y)
	fmt.Println("Correct time of flight (Seconds):", cTA,",",cTB)

}

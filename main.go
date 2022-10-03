package main

import (
	"fmt"
	"math"

	// "math"
	// "os"
	// "strconv"

	g "github.com/AllenDang/giu"
)

var (
	vel		float32 
	dist	float32
	h		float32

	thetaA	float64  
	thetaB	float64  
	maxH	float32

	timeA	float64
	timeB	float64

)

type coord struct {
	X	float64
	Y	float64
}

func (c coord) print(name string) {
	fmt.Print(name, ": (" , c.X , ", " , c.Y , ")\n" )
}


func loop() {
	var (
		lineA	[]float64 
		lineB	[]float64 
	)

	thetaA, thetaB = correctTheta(float64(vel), float64(dist), float64(h))
	maxH = float32(maxHeight(float64(vel), thetaB) * 1.1)

	timeA = timeOfFlight(float64(vel), thetaA, float64(h))
	timeB = timeOfFlight(float64(vel), thetaB, float64(h))

	for x := 0.0; x < float64(dist); x += 0.1 {
		lineA = append(lineA, vDispToX(float64(vel), x, thetaA))
		lineB = append(lineB, vDispToX(float64(vel), x, thetaB))
	}


	g.SingleWindow().Layout(
		g.Layout{
			g.Row(
				g.Column( 
					g.Label("Inputs"), 
					g.InputFloat(&vel).
						Size(float32(75)).
						Label("Velocity"), 
					g.InputFloat(&dist).
						Size(float32(75)).
						Label("Distance"), 
					g.InputFloat(&h).
						Size(float32(75)).
						Label("Height"),
					g.Label("Direct:"),
					g.Row( g.Column( g.Label("Angle"), g.Labelf("%.2f", thetaA) ),
						   g.Column( g.Label("Time"),  g.Labelf("%.2f", timeA)) ),
					g.Label("Loft:"),
					g.Row( g.Column( g.Label("Angle"), g.Labelf("%.2f", thetaB) ),
						   g.Column( g.Label("Time"),  g.Labelf("%.2f", timeB)) ),
				),
				g.Plot("Visual").AxisLimits(0, float64(dist * 1.05), 
				math.Min( float64(0 - maxH * 0.05), float64( h - maxH * 0.05 ) ),
					float64(maxH), g.ConditionAlways).Plots(
					g.PlotLine("Direct", lineA).XScale(0.1),
					g.PlotLine("Loft", lineB).XScale(0.1),
				),
			),
		},
	)


}

func main() {
	title := "Projectile Motion Calculator" 
	width, height := 800, 315
	wnd := g.NewMasterWindow(title, width, height, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop) 
}

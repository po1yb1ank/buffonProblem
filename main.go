package main

import (
	"fmt"
	"math"
	"math/rand"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)
func main() {
	var attempts int
	var t, l float64
	var events int
	fmt.Println("Input attempts number:")
	fmt.Scanf("%d\n", &attempts)
	fmt.Println("Input t len:")
	fmt.Scanf("%f\n", &t)
	fmt.Println("Input needle len:")
	fmt.Scanf("%f\n", &l)
	rand.Seed(rand.Int63())
	pts := make(plotter.XYs, attempts)
	for i := 1; i < attempts; i++{
		pts[i].X = float64(i)
		x := rand.Float64() * t / 2
		radian := float64(rand.Intn(180)) * (math.Pi/ 180.0)
		y := l/2 * math.Sin(radian)
		if x <= y{
			events++
		}
		if events != 0{
			pi := (2 * l * float64(i)) / (float64(events) * t)
			pts[i].Y = pi
		}

	}

	pi := (2 * l * float64(attempts)) / (float64(events) * t)
	fmt.Println(pi)
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Buffon problem"
	p.X.Label.Text = "attempt"
	p.Y.Label.Text = "Pi approximation"

	err = plotutil.AddLinePoints(p,
		"Pi on attempt", pts)
	p.Y.Min = 2
	p.Y.Max = 4
	p.Y.Dashes = []vg.Length{}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

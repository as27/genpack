package main

import (
	"flag"
	"time"

	"github.com/as27/genpack"
	"github.com/as27/gop5js"
)

var (
	flagMutationRate   = flag.Float64("mr", float64(0.01), "mutation rate")
	flagPopulation     = flag.Int("p", 30, "population")
	flagMaxGenerations = flag.Int("max", 100, "max generations")
)

var (
	startX  = 0.1
	startY  = 0.0
	targetX = 500.0
	targetY = 500.0
	frames  = 500
)

func fitness(d *genpack.DNS) float64 {
	m := newMover(startX, startY, d.Content)
	for i := 0; i <= frames; i++ {
		m.update()
	}
	dx := targetX - m.location.x
	dy := targetY - m.location.y
	return 1 / (dx*dx + dy*dy)
}

func main() {
	flag.Parse()
	mutationRate := *flagMutationRate
	population := *flagPopulation

	pop := genpack.CreateNewPopulation(
		population,
		frames,
		fitness,
		[]byte{0, 1, 2, 3, 4},
	)
	pop.CalcFitness()
	pop.Sort()
	counter := 0
	for {
		if counter == *flagMaxGenerations {
			pop.PrintN(10)
			break
		}
		counter++
		nextGen := pop.NextGeneration(mutationRate)
		pop = nextGen
		pop.CalcFitness()
		pop.Sort()
		pop.PrintN(10)
	}

	drawSetup()
	gop5js.Serve()
}

func drawSetup() {
	gop5js.Draw = draw
	gop5js.CanvasHeight = 600
	gop5js.CanvasWidth = 700
	gop5js.SleepPerFrame = time.Millisecond * time.Duration(100)
}

func draw() {
	gop5js.Background("127")

	//	m1.update()
	//	m1.draw()

}

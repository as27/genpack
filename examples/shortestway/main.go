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
	startX       = 250.1
	startY       = 250.1
	targetX      = 50.0
	targetY      = 50.0
	frames       = 50
	canvasHeight = 500
	canvasWidth  = 500
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

var pop *genpack.Population

func main() {
	flag.Float64Var(&startX, "sx", float64(250), "start X")
	flag.Float64Var(&startY, "sy", float64(250), "start Y")
	flag.Float64Var(&targetX, "tx", float64(10), "target X")
	flag.Float64Var(&targetY, "ty", float64(10), "target Y")
	flag.IntVar(&frames, "fr", 50, "frame number")
	flag.Parse()
	genpack.TimeSeed()
	drawSetup()
	go gop5js.Serve()
	population := *flagPopulation

	pop = genpack.CreateNewPopulation(
		population,
		frames,
		fitness,
		[]byte{0, 1, 2, 3, 4},
	)
	pop.CalcFitness()
	pop.Sort()
	var counter = 0

	for {
		if counter == *flagMaxGenerations {
			pop.PrintN(10)
			break
		}
		counter++
		createNextGeneration()
		drawPop()

	}

}
func createNextGeneration() {
	mutationRate := *flagMutationRate
	nextGen := pop.NextGeneration(mutationRate)
	pop = nextGen
	pop.CalcFitness()
	pop.Sort()
}

func drawSetup() {
	gop5js.Draw = draw
	gop5js.CanvasHeight = canvasHeight
	gop5js.CanvasWidth = canvasWidth
	gop5js.SleepPerFrame = time.Millisecond * time.Duration(100)

}

var movers []*Mover
var frameOver = make(chan bool)

func drawPop() {
	movers = []*Mover{}
	for _, d := range pop.DNSs {
		m := newMover(startX, startY, d.Content)
		movers = append(movers, m)
	}
	for i := 0; i < frames; i++ {
		<-frameOver
	}
}

func draw() {
	gop5js.Background("127")
	gop5js.Rect(targetX, targetY, 5, 5)
	for _, m := range movers {
		m.update()
		m.draw()
	}
	frameOver <- true
}

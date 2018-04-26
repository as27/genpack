package main

import (
	"flag"

	"github.com/as27/genpack"
	"github.com/as27/gop5js"
)

/*
Set the flags for the genetic algorithm
*/
var (
	flagMutationRate   = flag.Float64("mr", float64(0.01), "mutation rate")
	flagPopulation     = flag.Int("p", 30, "population")
	flagMaxGenerations = flag.Int("max", 100, "max generations")
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
var generationCounter = 0

var movers []*Mover
var frameOver = make(chan bool)

func main() {
	p5jsFlags()
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

	for {
		if generationCounter == *flagMaxGenerations {
			pop.PrintN(10)
			break
		}
		generationCounter++
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

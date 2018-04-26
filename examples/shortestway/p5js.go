package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/as27/gop5js"
)

var (
	startX       = 250.1
	startY       = 250.1
	targetX      = 50.0
	targetY      = 50.0
	frames       = 50
	canvasHeight = 480
	canvasWidth  = 852
)

func drawSetup() {
	gop5js.Draw = draw
	gop5js.CanvasHeight = canvasHeight
	gop5js.CanvasWidth = canvasWidth
	gop5js.SleepPerFrame = time.Millisecond * time.Duration(100)

}

func p5jsFlags() {
	flag.Float64Var(&startX, "sx", float64(150), "start X")
	flag.Float64Var(&startY, "sy", float64(150), "start Y")
	flag.Float64Var(&targetX, "tx", float64(350), "target X")
	flag.Float64Var(&targetY, "ty", float64(350), "target Y")
	flag.IntVar(&frames, "fr", 50, "frame number")
}

func draw() {
	gop5js.Background("127")
	gop5js.Text("Target", targetX+20, targetY)
	gop5js.Rect(targetX, targetY, 15, 15)
	gop5js.Text(fmt.Sprintf("Generation: %03d", generationCounter), 10, 50)
	for i, m := range movers {
		m.update()
		m.size = 5 + (float64(len(movers)) / float64((i + 1)))

		m.draw()
	}
	frameOver <- true
}

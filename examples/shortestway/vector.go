package main

import (
	"github.com/as27/gop5js"
)

var (
	vdown  = Vector{0, 0.1}
	vup    = Vector{0, -0.1}
	vleft  = Vector{-0.1, 0}
	vright = Vector{0.1, 0}
)

type Vector struct {
	x, y float64
}

func (v *Vector) add(v2 *Vector) {
	v.x = v.x + v2.x
	v.y = v.y + v2.y
}

type Mover struct {
	location     *Vector
	velocity     *Vector
	acceleration *Vector
	dna          []byte
	pointer      int
}

func newMover(x, y float64, dna []byte) *Mover {
	m := Mover{dna: dna}
	m.pointer = 0
	m.location = &Vector{x, y}
	m.velocity = &Vector{0, 0}
	m.acceleration = &Vector{0, 0}
	return &m
}

func (m *Mover) accelerate(v *Vector) {
	m.acceleration.add(v)
}

func (m *Mover) update() {
	if m.pointer >= len(m.dna) {
		m.pointer = 0
	}
	var v *Vector
	switch m.dna[m.pointer] {
	case 1:
		v = &vup
	case 2:
		v = &vdown
	case 3:
		v = &vleft
	case 4:
		v = &vright
	default:
		v = &Vector{0, 0}
	}
	m.accelerate(v)
	m.pointer++
	m.velocity.add(m.acceleration)
	m.location.add(m.velocity)
}

func (m *Mover) draw() {
	gop5js.Ellipse(m.location.x, m.location.y, 15, 15)
}

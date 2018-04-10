package genpack

import (
	"fmt"
	"math/rand"
)

type Fitnesser interface {
	Fitness() float64
}

type DNS struct {
	fitnesser Fitnesser
	Content   []byte
	fitness   float64
}

func NewDNS(b []byte, f Fitnesser) *DNS {
	return &DNS{
		fitnesser: f,
		Content:   b,
	}
}

func NewRandomDNS(length int, allowedBytes []byte, f Fitnesser) *DNS {
	var b []byte
	for i := 0; i < length; i++ {
		b = append(b, randByte(allowedBytes))
	}
	return NewDNS(b, f)
}

func (d *DNS) mutate(mutationRate float64, allowedBytes []byte) {
	for i := range d.Content {
		if rand.Float64() <= mutationRate {
			d.Content[i] = randByte(allowedBytes)
		}
	}
}

func randByte(allowedBytes []byte) byte {
	return allowedBytes[rand.Intn(len(allowedBytes))]
}

func (d *DNS) Reproduce(father *DNS) (*DNS, *DNS) {
	childDNS1 := make([]byte, len(d.Content))
	childDNS2 := make([]byte, len(d.Content))
	for i := range childDNS1 {
		if i%2 == 0 {
			childDNS1[i] = father.Content[i]
			childDNS2[i] = d.Content[i]
		} else {
			childDNS1[i] = d.Content[i]
			childDNS2[i] = father.Content[i]
		}
	}

	return NewDNS(childDNS1, d.fitnesser), NewDNS(childDNS2, d.fitnesser)
}

func (d *DNS) String() string {
	return fmt.Sprintf("%s", string(d.Content))
}

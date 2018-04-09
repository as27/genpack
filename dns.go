package genpack

import (
	"fmt"
	"math/rand"
)

type DNS struct {
	content []byte
	Fitness float64
}

func NewDNS(b []byte) *DNS {
	return &DNS{
		content: b,
	}
}

	var b []byte
	for i := 0; i < length; i++ {
		b = append(b, randByte(allowedBytes))
	}
	return NewDNS(b)
}

func (d *DNS) mutate(mutationRate float64, allowedBytes []byte) {
	for i := range d.content {
		if rand.Float64() <= mutationRate {
			d.content[i] = randByte(allowedBytes)
		}
	}
}

func randByte(allowedBytes []byte) byte {
	return allowedBytes[rand.Intn(len(allowedBytes))]
}

func (d *DNS) Reproduce(father *DNS) (*DNS, *DNS) {
	childDNS1 := make([]byte, len(d.content))
	childDNS2 := make([]byte, len(d.content))
	for i := range childDNS1 {
		if i%2 == 0 {
			childDNS1[i] = father.content[i]
			childDNS2[i] = d.content[i]
		} else {
			childDNS1[i] = d.content[i]
			childDNS2[i] = father.content[i]
		}
	}

	return NewDNS(childDNS1), NewDNS(childDNS2)
}

func (d *DNS) String() string {
	return fmt.Sprintf("%s", string(d.content))
}

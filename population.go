package genpack

import (
	"fmt"
	"math/rand"
	"sort"
)

// Population contains a population for the genetic algorithm
type Population struct {
	dnss         []*DNS
	fitnessSum   float64
	fitnessFunc  func(*DNS) float64
	allowedBytes []byte
}

// CreateNewPopulation generates a population
func CreateNewPopulation(
	popSize int,
	dnsLength int,
	fitnessFunc func(*DNS) float64,
	allowedBytes []byte,
) *Population {
	dnss := make([]*DNS, popSize)
	for i := 0; i < popSize; i++ {
		dnss[i] = NewRandomDNS(dnsLength, allowedBytes)
	}
	return &Population{
		dnss:         dnss,
		fitnessFunc:  fitnessFunc,
		allowedBytes: allowedBytes,
	}
}

// CalcFitness calculates the fitness for all dns of the population
func (p *Population) CalcFitness() {
	for _, d := range p.dnss {
		go func(d *DNS) {
			d.Fitness = p.fitnessFunc(d)
		}(d)

	}
}

func (p *Population) NextGeneration(mutationRate float64) *Population {
	ng := Population{}
	ng.fitnessFunc = p.fitnessFunc
	dnss := make([]*DNS, p.Size())
	for i := 0; i < p.Size(); i++ {
		dnsMum := p.PickDNS()
		dnsDad := p.PickDNS()
		child, _ := dnsMum.Reproduce(dnsDad)
		child.mutate(mutationRate, p.allowedBytes)

		dnss[i] = child
	}
	ng.dnss = dnss
	return &ng
}

func (p *Population) PickDNS() *DNS {
	if p.fitnessSum == 0 {
		var fitnessSum float64
		for _, d := range p.dnss {
			fitnessSum = fitnessSum + d.Fitness
		}
		p.fitnessSum = fitnessSum
	}

	r := rand.Float64() * float64(p.fitnessSum)
	fitMin := 0.0
	fitMax := 0.0
	for _, d := range p.dnss {
		fitMax = fitMin + d.Fitness
		if fitMin <= r && r <= fitMax {
			return d
		}
		fitMin = fitMax
	}
	fmt.Println("----->", r, p.fitnessSum)
	return p.dnss[0]
}

func (p *Population) Size() int {
	return len(p.dnss)
}

// Sort the population according to the fitness of the dnss higher
// fitness is sorted first
func (p *Population) Sort() {
	sort.Slice(p.dnss, func(i, j int) bool {
		return p.dnss[i].Fitness > p.dnss[j].Fitness
	})
}

func (p *Population) PrintN(n int) {
	for i, dns := range p.dnss {
		if i == n {
			break
		}
		if dns == nil {
			continue
		}
		fmt.Println(
			dns,
			dns.Fitness)
	}

	fmt.Println("----------------------")
}

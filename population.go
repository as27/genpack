package genpack

import (
	"fmt"
	"math/rand"
	"sort"
)

// Population contains a population for the genetic algorithm
type Population struct {
	DNSs         []*DNS
	fitnessSum   float64
	fitnessFunc  func(*DNS) float64
	allowedBytes []byte
}

// CreateNewPopulation generates a population. All compulsory elements
// are input parameters.
//   popSize:      size of the population
//   dnsLength:    length of the dns string
//   fitnessFunc:  function which calculates the fitness
//   allowedBytes: slice of the allowed bytes of the DNS string
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
		DNSs:         dnss,
		fitnessFunc:  fitnessFunc,
		allowedBytes: allowedBytes,
	}
}

// CalcFitness calculates the fitness for all dns of the population
func (p *Population) CalcFitness() {
	for _, d := range p.DNSs {
		go func(d *DNS) {
			d.fitness = p.fitnessFunc(d)
		}(d)

	}
}

// NextGeneration generates a new population
func (p *Population) NextGeneration(mutationRate float64) *Population {
	ng := Population{}
	ng.allowedBytes = p.allowedBytes
	ng.fitnessFunc = p.fitnessFunc
	dnss := make([]*DNS, p.Size())
	for i := 0; i < p.Size(); i++ {
		dnsMum := p.pickDNS()
		dnsDad := p.pickDNS()
		child, _ := dnsMum.Reproduce(dnsDad)
		child.mutate(mutationRate, p.allowedBytes)

		dnss[i] = child
	}
	ng.DNSs = dnss
	return &ng
}

func (p *Population) pickDNS() *DNS {
	if p.fitnessSum == 0 {
		var fitnessSum float64
		for _, d := range p.DNSs {
			fitnessSum = fitnessSum + d.fitness
		}
		p.fitnessSum = fitnessSum
	}

	r := rand.Float64() * float64(p.fitnessSum)
	fitMin := 0.0
	fitMax := 0.0
	for _, d := range p.DNSs {
		fitMax = fitMin + d.fitness
		if fitMin <= r && r <= fitMax {
			return d
		}
		fitMin = fitMax
	}
	fmt.Println("----->", r, p.fitnessSum)
	return p.DNSs[0]
}

// Size of the population. That means the number of DNS inside that population
func (p *Population) Size() int {
	return len(p.DNSs)
}

// Sort the population according to the fitness of the dnss higher
// fitness is sorted first
func (p *Population) Sort() {
	sort.Slice(p.DNSs, func(i, j int) bool {
		return p.DNSs[i].fitness > p.DNSs[j].fitness
	})
}

func (p *Population) PrintN(n int) {
	for i, dns := range p.DNSs {
		if i == n {
			break
		}
		if dns == nil {
			continue
		}
		fmt.Println(
			dns,
			dns.fitness)
	}

	fmt.Println("----------------------")
}

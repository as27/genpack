package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/as27/genpack"
)

func init() {
	genpack.Seed(time.Now().UnixNano())
}

var (
	flagMutationRate   = flag.Float64("mr", float64(0.01), "mutation rate")
	flagPopulation     = flag.Int("p", 800, "population")
	flagTargetText     = ""
	flagMaxGenerations = flag.Int("max", 5000, "max generations")
)

type Text struct {
	txt string
}

func NewText() genpack.DNSFitnesser {
	return &Text{}
}

func (t *Text) LoadDNS(dns *genpack.DNS) {
	t.txt = string(dns.Content)
}

func (t *Text) Fitness() float64 {
	match := 0
	for i, b := range flagTargetText {
		if b == rune(t.txt[i]) {
			match++
		}
	}
	//fmt.Println(t.txt, flagTargetText, match)
	return float64(match)
}

func main() {
	flag.StringVar(&flagTargetText, "t", "this is the secret text", "target text")
	flag.Parse()
	start := time.Now()
	targetText := flagTargetText
	mutationRate := *flagMutationRate
	population := *flagPopulation
	fmt.Println(targetText, mutationRate, population)
	pop := genpack.CreateNewPopulation(
		population,
		len(targetText),
		NewText,
		[]byte("abcde"),
	)

	pop.CalcFitness()
	pop.Sort()

	counter := 0
	for {
		if counter%20 == 0 {
			go pop.PrintN(5)
		}

		//wait()
		bestMatch := string(pop.DNSs[0].Content)
		//fmt.Println(pop.Fitness(pop.dnss[0]), bestMatch, pop.dnss[0].content, pop.dnss[1].content, pop.dnss[2].content)
		if bestMatch == targetText {
			fmt.Println(targetText, mutationRate, population)
			fmt.Println("Counter", counter, bestMatch)
			end := time.Now()
			fmt.Println("Time:", end.Sub(start))
			break
		}
		if counter == *flagMaxGenerations {
			pop.PrintN(10)
			break
		}

		counter++
		nextGen := pop.NextGeneration(mutationRate)

		pop = nextGen
		pop.CalcFitness()
		pop.Sort()
	}
}
func wait() {
	var userIn string
	fmt.Println("Press ENTER to continue")
	fmt.Println("Press q to quit")
	fmt.Scanln(&userIn)
	if userIn == "q" {
		os.Exit(4)
	}
}

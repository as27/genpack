package main

import (
	"fmt"

	"github.com/as27/genpack"
)

var targetX = 600
var targetY = 500

func fitness(dns *genpack.DNS) float64 {
	return 0.0
}

func main() {
	pop := genpack.CreateNewPopulation(20, 100, fitness, []byte{0, 1, 2, 3, 4})
	fmt.Println(pop)
}

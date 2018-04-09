/*
Package genpack implements genetic algorithm in Go
*/
package genpack

import "math/rand"

// Seed passes the value direct to rand.Seed(). That the pseudo random numbers
// are working you need to pass a seed here. If you already set a seed for the
// rand package you don't have to call this function.
func Seed(seed int64) {
	rand.Seed(seed)
}

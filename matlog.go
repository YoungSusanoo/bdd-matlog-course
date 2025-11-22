package main

import (
	"fmt"

	"github.com/dalzilio/rudd"
)

const (
	objects   = 9
	props     = 4
	propsBits = 4
)

func main() {
	bdd, err := rudd.New(6, rudd.Nodesize(10000), rudd.Cachesize(5000))
	v := make([][][]rudd.BDD, objects)
	v[0] = make([][]rudd.BDD, props)
	v[0][0] = make([]rudd.BDD, propsBits)

	fmt.Print(err)
}

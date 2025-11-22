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
	if err != nil {
		fmt.Println(err)
	}
	
	vars := make([][][]rudd.Node, objects)
	for i := range objects {
		vars[i] = make([][]rudd.Node, props)
		for j := range props {
			vars[i][j] = make([]rudd.Node, propsBits)
			for k := range propsBits {
				vars[i][j][k] = bdd.Ithvar(objects*i + props*j + k)
			}
		}
	}
}

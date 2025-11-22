package buildtree

import (
	"math"

	"github.com/dalzilio/rudd"
)

func countBits(props int) int {
	propsInt := int(props)
	return int(math.Ceil(math.Log(float64(propsInt))))
}

func Build(bdd rudd.BDD, objects, props int) (vars [][][]rudd.Node) {
	vars = make([][][]rudd.Node, objects)
	for i := range objects {
		vars[i] = make([][]rudd.Node, props)
		for j := range props {
			vars[i][j] = make([]rudd.Node, countBits(props))
			for k := range countBits(props) {
				bit := j >> k
				if bit == 1 {
					vars[i][j][k] = bdd.Ithvar(i*objects + j*props + k)
				} else {
					vars[i][j][k] = bdd.NIthvar(i*objects + j*props + k)
				}
			}
		}
	}
	return
}

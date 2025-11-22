package conditions

import (
	"github.com/dalzilio/rudd"
)

type Conditions struct {
	Bdd            *rudd.BDD
	Objects, Props int
}

func (c Conditions) AddSpecificProp(vars [][][]rudd.Node, object, prop, propVal int) rudd.Node {
	return vars[object][prop][propVal]
}

func (c Conditions) AddPropsRelation(vars [][][]rudd.Node, prop1, propVal1, prop2, propVal2 int) (newtree rudd.Node) {
	newtree = c.Bdd.True()
	for i := range c.Objects {
		c.Bdd.And(c.Bdd.Not(vars[i][prop1][propVal1]), c.Bdd.Not(vars[i][prop2][propVal2]))
	}
	return
}

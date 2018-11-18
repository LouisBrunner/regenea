package core

import (
	"github.com/LouisBrunner/regenea/genea"
)

type Processor interface {
	ProcessPerson(person *genea.Person)
	ProcessUnion(union *genea.Union)
	Finish()
}

func ProcessTree(tree *genea.Tree, processors []Processor) {
	for _, person := range tree.People {
		for _, p := range processors {
			p.ProcessPerson(person)
		}
	}
	tree.IterateOverUnion(func(union *genea.Union) {
		for _, p := range processors {
			p.ProcessUnion(union)
		}
	})
	for _, p := range processors {
		p.Finish()
	}
}

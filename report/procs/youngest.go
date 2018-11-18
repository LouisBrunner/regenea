package procs

import (
	"github.com/LouisBrunner/regenea/genea"
)

type Youngest struct {
}

func (p *Youngest) ProcessPerson(person *genea.Person) {
}

func (p *Youngest) ProcessUnion(union *genea.Union) {
}

func (p *Youngest) Finish() {
}

func (p *Youngest) Output() (string, interface{}) {
	return CategoryGeneral, nil
}

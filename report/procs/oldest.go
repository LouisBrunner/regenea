package procs

import (
	"github.com/LouisBrunner/regenea/genea"
)

type Oldest struct {
}

func (p *Oldest) ProcessPerson(person *genea.Person) {
}

func (p *Oldest) ProcessUnion(union *genea.Union) {
}

func (p *Oldest) Finish() {
}

func (p *Oldest) Output() (string, interface{}) {
	return CategoryGeneral, nil
}

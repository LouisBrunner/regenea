package procs

import (
	"fmt"

	"github.com/LouisBrunner/regenea/genea"

	"github.com/dustin/go-humanize"
)

type buRelation struct {
	uncle  *genea.Person
	nephew *genea.Person
}

// TODO: orphaned at birth?
// TODO: incest calculation?

type Weird struct {
	posthumous   []*genea.Person
	fastWeddings []*genea.Person
	babyUncles   []*buRelation
}

func (p *Weird) ProcessPerson(person *genea.Person) {
	if genea.EventIsValid(&person.Birth) {
		if person.Father != nil && genea.EventIsValid(person.Father.Death) {
			if person.Birth.Date.After(person.Father.Death.Date) {
				p.posthumous = append(p.posthumous, person)
			}
		}

		if person.Family != nil && genea.EventIsValid(&person.Family.Begin) {
			if person.Birth.Date.Add(-280 * humanize.Day).Before(person.Family.Begin.Date) {
				p.fastWeddings = append(p.fastWeddings, person)
			}
		}

		var parentNephew *genea.Person
		for _, sibling := range person.Siblings {
			for _, nephew := range sibling.Children {
				if genea.EventIsValid(&nephew.Birth) && person.Birth.Date.After(nephew.Birth.Date) {
					parentNephew = nephew
					break
				}
			}
		}

		if parentNephew != nil {
			p.babyUncles = append(p.babyUncles, &buRelation{
				uncle:  person,
				nephew: parentNephew,
			})
		}
	}
}

func (p *Weird) ProcessUnion(union *genea.Union) {
}

func (p *Weird) Finish() {
}

func (p *Weird) Output() (string, StringMap) {
	posthumous := make([]string, len(p.posthumous))
	for i, ph := range p.posthumous {
		_, pronoun, _ := ph.Pronouns()
		posthumous[i] = fmt.Sprintf(
			"%s was born on the %s but %s father died the %s (%s)",
			ph.Name(),
			ph.Birth.FormatDate(),
			pronoun,
			ph.Father.Death.FormatDate(),
			humanize.CustomRelTime(ph.Birth.Date, ph.Father.Death.Date, "earlier", "later", myMagnitudes),
		)
	}

	fastWeddings := make([]string, len(p.fastWeddings))
	for i, fw := range p.fastWeddings {
		_, pronoun, _ := fw.Pronouns()
		fastWeddings[i] = fmt.Sprintf(
			"%s was born on the %s but %s parents got married on the %s (%s)",
			fw.Name(),
			fw.Birth.FormatDate(),
			pronoun,
			fw.Family.Begin.FormatDate(),
			humanize.CustomRelTime(fw.Birth.Date, fw.Family.Begin.Date, "earlier", "later", myMagnitudes),
		)
	}

	babyUncles := make([]string, len(p.babyUncles))
	for i, bu := range p.babyUncles {
		pronounUncle, _, _ := bu.uncle.Pronouns()
		babyUncles[i] = fmt.Sprintf(
			"%s was born on the %s but %s was already the %s of %s who was born on the %s (%s)",
			bu.uncle.Name(),
			bu.uncle.Birth.FormatDate(),
			pronounUncle,
			"uncle/aunt", // TODO: need a function to describe relation
			bu.nephew.Name(),
			bu.nephew.Birth.FormatDate(),
			humanize.CustomRelTime(bu.uncle.Birth.Date, bu.nephew.Birth.Date, "earlier", "later", myMagnitudes),
		)
	}

	return categoryGeneral, StringMap{
		"Posthumous children":                   posthumous,
		"Children conceived before the wedding": fastWeddings,
		"Baby uncles/aunts":                     babyUncles,
	}
}

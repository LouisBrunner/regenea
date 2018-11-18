package procs

import (
	"fmt"
	"time"

	"github.com/LouisBrunner/regenea/genea"

	"github.com/dustin/go-humanize"
)

type average struct {
	sum   float32
	count uint64
}

func (avg *average) calculate() float32 {
	if avg.count == 0 {
		return 0
	}
	return avg.sum / float32(avg.count)
}

type personAge struct {
	person *genea.Person
	age    uint64
}

type unionAge struct {
	union *genea.Union
	age   uint64
}

// TODO: every stat living vs dead
// TODO: old/young/avg when married
// TODO: old/young/avg when first kid
// TODO: old/young/avg when last kid
// TODO: old/young/avg when avg kid?
// TODO: old/young/avg when orphaned
// TODO: old/young/avg when widowed

type Ages struct {
	youngest      personAge
	youngestMan   personAge
	youngestWoman personAge
	youngestOther personAge

	oldest      personAge
	oldestMan   personAge
	oldestWoman personAge
	oldestOther personAge

	meanAge      average
	meanAgeMen   average
	meanAgeWomen average
	meanAgeOther average

	oldestUnion   unionAge
	youngestUnion unionAge
	meanAgeUnion  average
}

func (p *Ages) ProcessPerson(person *genea.Person) {
	if !genea.EventIsValid(&person.Birth) || (person.Death != nil && !genea.EventIsValid(person.Death)) {
		return
	}

	death := time.Now()
	if person.Death != nil {
		death = person.Death.Date
	}
	age := uint64(death.Unix() - person.Birth.Date.Unix())
	agef := float32(age) / float32(60) / float32(60) / float32(24) / float32(timeYear)

	if p.oldest.person == nil || p.oldest.age < age {
		p.oldest.person = person
		p.oldest.age = age
	}
	if p.youngest.person == nil || p.youngest.age > age {
		p.youngest.person = person
		p.youngest.age = age
	}
	p.meanAge.sum += agef
	p.meanAge.count += 1

	switch person.Sex {
	case genea.SexMale:
		if p.oldestMan.person == nil || p.oldestMan.age < age {
			p.oldestMan.person = person
			p.oldestMan.age = age
		}
		if p.youngestMan.person == nil || p.youngestMan.age > age {
			p.youngestMan.person = person
			p.youngestMan.age = age
		}
		p.meanAgeMen.sum += agef
		p.meanAgeMen.count += 1
	case genea.SexFemale:
		if p.oldestWoman.person == nil || p.oldestWoman.age < age {
			p.oldestWoman.person = person
			p.oldestWoman.age = age
		}
		if p.youngestWoman.person == nil || p.youngestWoman.age > age {
			p.youngestWoman.person = person
			p.youngestWoman.age = age
		}
		p.meanAgeWomen.sum += agef
		p.meanAgeWomen.count += 1
	default:
		if p.oldestOther.person == nil || p.oldestOther.age < age {
			p.oldestOther.person = person
			p.oldestOther.age = age
		}
		if p.youngestOther.person == nil || p.youngestOther.age > age {
			p.youngestOther.person = person
			p.youngestOther.age = age
		}
		p.meanAgeOther.sum += agef
		p.meanAgeOther.count += 1
	}
}

func (p *Ages) ProcessUnion(union *genea.Union) {
	if !genea.EventIsValid(&union.Begin) || (union.Person1 == nil && union.Person2 == nil && !genea.EventIsValid(union.End)) {
		return
	}

	end := getUnionEnd(union)
	if end.IsZero() {
		return
	}
	age := uint64(end.Unix() - union.Begin.Date.Unix())
	agef := float32(age) / float32(60) / float32(60) / float32(24) / float32(timeYear)

	p.meanAgeUnion.sum += agef
	p.meanAgeUnion.count += 1

	if p.oldestUnion.union == nil || p.oldestUnion.age < age {
		p.oldestUnion.union = union
		p.oldestUnion.age = age
	}
	if p.youngestUnion.union == nil || p.youngestUnion.age > age {
		p.youngestUnion.union = union
		p.youngestUnion.age = age
	}
}

func (p *Ages) Finish() {
}

func getUnionEnd(union *genea.Union) time.Time {
	now := time.Now()
	if genea.EventIsValid(union.End) {
		return union.End.Date
	}
	person1d := now
	person2d := now
	if union.Person1 != nil {
		if genea.EventIsValid(union.Person1.Death) {
			person1d = union.Person1.Death.Date
		} else if union.Person1.Death != nil {
			return time.Time{}
		}
	}
	if union.Person2 != nil && genea.EventIsValid(union.Person2.Death) {
		if genea.EventIsValid(union.Person2.Death) {
			person2d = union.Person2.Death.Date
		} else if union.Person2.Death != nil {
			return time.Time{}
		}
	}
	if person1d.Before(person2d) {
		return person1d
	}
	return person2d
}

func formatPerson(pa personAge) string {
	person := pa.person
	if person == nil {
		return "none"
	}

	death := time.Now()
	deathStr := "Present"
	verb := "is"
	if person.Death != nil {
		death = person.Death.Date
		deathStr = person.Death.FormatDate()
		verb = "was"
	}
	return fmt.Sprintf(
		"%s %s %s (%s - %s)",
		person.Name(),
		verb,
		humanize.CustomRelTime(person.Birth.Date, death, "old", "ERROR", myMagnitudes),
		person.Birth.FormatDate(),
		deathStr,
	)
}

func formatUnion(ua unionAge) string {
	union := ua.union
	if union == nil {
		return "none"
	}

	person1 := "Unknown"
	if union.Person1 != nil {
		person1 = union.Person1.Name()
	}
	person2 := "Unknown"
	if union.Person2 != nil {
		person2 = union.Person2.Name()
	}
	end := getUnionEnd(union)
	endStr := "Present"
	verb := "is"
	if time.Since(end) > humanize.Day {
		endStr = genea.FormatDate(end)
		verb = "was"
	}
	return fmt.Sprintf(
		"Union between %s and %s %s %s (%s - %s)",
		person1, person2,
		verb,
		humanize.CustomRelTime(union.Begin.Date, end, "old", "ERROR", myMagnitudes),
		union.Begin.FormatDate(),
		endStr,
	)
}

func (p *Ages) Output() (string, StringMap) {
	return categoryGeneral, StringMap{
		"Youngest": StringMap{
			"Person": formatPerson(p.youngest),
			"Man":    formatPerson(p.youngestMan),
			"Woman":  formatPerson(p.youngestWoman),
			"Other":  formatPerson(p.youngestOther),
		},
		"Oldest": StringMap{
			"Person": formatPerson(p.oldest),
			"Man":    formatPerson(p.oldestMan),
			"Woman":  formatPerson(p.oldestWoman),
			"Other":  formatPerson(p.oldestOther),
		},
		"Mean": StringMap{
			"All":   p.meanAge.calculate(),
			"Men":   p.meanAgeMen.calculate(),
			"Women": p.meanAgeWomen.calculate(),
			"Other": p.meanAgeOther.calculate(),
		},
		"Union": StringMap{
			"Longest":  formatUnion(p.oldestUnion),
			"Shortest": formatUnion(p.youngestUnion),
			"Mean age": p.meanAgeUnion.calculate(),
		},
	}
}

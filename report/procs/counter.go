package procs

import (
	"github.com/LouisBrunner/regenea/genea"
)

type Counter struct {
	personCount uint64
	manCount    uint64
	womanCount  uint64
	otherCount  uint64

	unionCount   uint64
	weddingCount uint64
	civilCount   uint64
	sameSexCount uint64
	diffSexCount uint64
}

func (p *Counter) ProcessPerson(person *genea.Person) {
	p.personCount++
	switch person.Sex {
	case genea.SexMale:
		p.manCount++
	case genea.SexFemale:
		p.womanCount++
	default:
		p.otherCount++
	}
}

func (p *Counter) ProcessUnion(union *genea.Union) {
	p.unionCount++
	switch union.Kind {
	case genea.UnionWedding:
		p.weddingCount++
	case genea.UnionCivil:
		p.civilCount++
	}
	if union.Person1 != nil && union.Person2 != nil {
		if union.Person1.Sex == union.Person2.Sex {
			p.sameSexCount++
		} else {
			p.diffSexCount++
		}
	}
}

func (p *Counter) Finish() {
}

func (p *Counter) Output() (string, StringMap) {
	totalpf := float32(p.personCount)
	if p.personCount == 0 {
		totalpf = 1
	}
	totaluf := float32(p.unionCount)
	if p.unionCount == 0 {
		totaluf = 1
	}
	unknownSexUnion := p.unionCount - (p.sameSexCount + p.diffSexCount)

	return categoryGeneral, StringMap{
		"People count": StringMap{
			"Total": p.personCount,
			"Men":   p.manCount,
			"Women": p.womanCount,
			"Other": p.otherCount,
		},
		"People ratio": StringMap{
			"Men":   float32(p.manCount) / totalpf,
			"Women": float32(p.womanCount) / totalpf,
			"Other": float32(p.otherCount) / totalpf,
		},
		"Unions count": StringMap{
			"Total": p.unionCount,
			"Kind": StringMap{
				"Wedding": p.weddingCount,
				"Civil":   p.civilCount,
			},
			"Sex": StringMap{
				"Same":      p.sameSexCount,
				"Different": p.diffSexCount,
				"Unknown":   unknownSexUnion,
			},
		},
		"Unions ratio": StringMap{
			"Kind": StringMap{
				"Wedding": float32(p.weddingCount) / totaluf,
				"Civil":   float32(p.civilCount) / totaluf,
			},
			"Sex": StringMap{
				"Same":      float32(p.sameSexCount) / totaluf,
				"Different": float32(p.diffSexCount) / totaluf,
				"Unknown":   float32(unknownSexUnion) / totaluf,
			},
		},
	}
}

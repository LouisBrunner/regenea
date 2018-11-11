package genea

import (
	"time"

	"github.com/LouisBrunner/regenea/genea/json"
)

type EventCore struct {
	Date     time.Time
	Location string
	Source   string
}

type Event struct {
	EventCore
	Title string
}

type UnionKind int

const (
	UnionWedding UnionKind = iota
	UnionCivil
)

type Union struct {
	Kind     UnionKind
	Person1  *Person
	Person2  *Person
	Begin    time.Time
	End      time.Time
	Issues   []*Person
	Comments string
}

type Sex int

const (
	SexMale Sex = iota
	SexFemale
	SexOther
)

type Person struct {
	ID       string
	Names    json.Names
	Sex      Sex
	Birth    EventCore
	Death    *EventCore
	Events   []Event
	Comments string

	Father   *Person
	Mother   *Person
	Partners []*Union
	Children []*Person
	Siblings []*Person

	flags struct {
		issuesFound bool
	}
}

type Tree struct {
	Comments string
	People   []Person
	ByID     map[string]*Person
}

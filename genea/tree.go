package genea

import (
	"time"

	"github.com/LouisBrunner/regenea/genea/json"
)

// TODO: add way to get subset of a tree (all parents, siblings (child included) and children)

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
	Begin    EventCore
	End      *EventCore
	Issues   []*Person
	Comments string

	flags struct {
		serialized bool
	}
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

	Family   *Union
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
	People   []*Person
	ByID     map[string]*Person
}

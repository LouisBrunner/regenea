package genea

import (
	"bytes"
	"time"

	"github.com/LouisBrunner/regenea/genea/json"
)

func EventIsValid(event *EventCore) bool {
	return event != nil && event.IsValid()
}

func (event *EventCore) IsValid() bool {
	return !event.Date.IsZero()
}

func ConstructName(names *json.Names) string {
	buf := &bytes.Buffer{}
	if names.First != "" {
		buf.WriteString(names.First)
	}
	if names.Middle != "" {
		if buf.Len() != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(names.Middle)
	}
	if names.Last != "" {
		if buf.Len() != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(names.Last)
	}
	if names.Alternative != "" {
		if buf.Len() != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("aka ")
		buf.WriteString(names.Alternative)
	}
	return buf.String()
}

func (p *Person) Name() string {
	return ConstructName(&p.Names)
}

func (p *Person) Pronouns() (noun string, accusatif string, possessive string) {
	switch p.Sex {
	case SexMale:
		return "he", "him", "his"
	case SexFemale:
		return "she", "her", "her"
	default:
		return "they", "their", "their"
	}
}

func (e *EventCore) FormatDate() string {
	return FormatDate(e.Date)
}

func FormatDate(tm time.Time) string {
	return tm.Format(json.DateFormat)
}

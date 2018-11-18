package genea

import (
	"bytes"
	"time"

	"github.com/LouisBrunner/regenea/genea/json"
)

func EventIsValid(event *EventCore) bool {
	return event != nil && !event.Date.IsZero()
}

func (p *Person) Name() string {
	buf := &bytes.Buffer{}
	if p.Names.First != "" {
		buf.WriteString(p.Names.First)
	}
	if p.Names.Middle != "" {
		if buf.Len() != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(p.Names.Middle)
	}
	if p.Names.Last != "" {
		if buf.Len() != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(p.Names.Last)
	}
	if p.Names.Alternative != "" {
		if buf.Len() != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("aka ")
		buf.WriteString(p.Names.Alternative)
	}
	return buf.String()
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

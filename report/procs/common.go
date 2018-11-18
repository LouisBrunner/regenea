package procs

import (
	"github.com/LouisBrunner/regenea/core"
)

const (
	CategoryGeneral = "General"
	CategoryStats   = "General"
)

type Processor interface {
	core.Processor

	Output() (string, interface{})
}

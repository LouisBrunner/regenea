package procs

import (
	"math"
	"time"

	"github.com/LouisBrunner/regenea/core"

	"github.com/dustin/go-humanize"
)

const (
	categoryGeneral = "General"
)

type Processor interface {
	core.Processor

	Output() (string, StringMap)
}

const timeYear = 365

type StringMap map[string]interface{}

var myMagnitudes = []humanize.RelTimeMagnitude{
	{time.Second, "now", time.Second},
	{2 * time.Second, "1 second %s", 1},
	{time.Minute, "%d seconds %s", time.Second},
	{2 * time.Minute, "1 minute %s", 1},
	{time.Hour, "%d minutes %s", time.Minute},
	{2 * time.Hour, "1 hour %s", 1},
	{humanize.Day, "%d hours %s", time.Hour},
	{2 * humanize.Day, "1 day %s", 1},
	{humanize.Week, "%d days %s", humanize.Day},
	{2 * humanize.Week, "1 week %s", 1},
	{humanize.Month, "%d weeks %s", humanize.Week},
	{2 * humanize.Month, "1 month %s", 1},
	{humanize.Year, "%d months %s", humanize.Month},
	{18 * humanize.Month, "1 year %s", 1},
	{2 * humanize.Year, "2 years %s", 1},
	{math.MaxInt64, "%d years %s", humanize.Year},
}

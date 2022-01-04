package parser

import (
	"regexp"
)

var (
	freshnessUnits = `decade|year|month|week|day|hour|minute|second`
)

var (
	rxNonWord   = regexp.MustCompile(`\W`)
	rxDayNumber = regexp.MustCompile(`2.{0,3}`)
	rxBraces    = regexp.MustCompile(`[{}()<>\[\]]`)
	rxTimestamp = regexp.MustCompile(`^(\d{10})(\d{3})?(\d{3})?(?:\.|\s|$)`)

	rxIn                = regexp.MustCompile(`(?i)\bin\b`)
	rxAgo               = regexp.MustCompile(`(?i)\bago\b`)
	rxInAgo             = regexp.MustCompile(`(?i)\b(?:ago|in)\b`)
	rxFreshnessPattern  = regexp.MustCompile(`(?i)(\d+)\s*(` + freshnessUnits + `)\b`)
	rxFreshnessSkipWord = regexp.MustCompile(`(?i)^(?:` + freshnessUnits + `|ago|in|\d+|:|[ap]m)`)
)
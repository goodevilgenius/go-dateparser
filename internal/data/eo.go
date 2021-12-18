// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var eo_Locale = merge(nil, LocaleData{
	Name:      "eo",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"yesterday":   "1 day ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"septembro":   "september",
		"decembro":    "december",
		"februaro":    "february",
		"vendredo":    "friday",
		"tomorrow":    "in 1 day",
		"novembro":    "november",
		"merkredo":    "wednesday",
		"augusto":     "august",
		"januaro":     "january",
		"oktobro":     "october",
		"dimanco":     "sunday",
		"aprilo":      "april",
		"minute":      "minute",
		"sabato":      "saturday",
		"second":      "second",
		"today":       "0 day ago",
		"julio":       "july",
		"junio":       "june",
		"marto":       "march",
		"lundo":       "monday",
		"month":       "month",
		"jaudo":       "thursday",
		"mardo":       "tuesday",
		"hour":        "hour",
		"majo":        "may",
		"week":        "week",
		"year":        "year",
		"now":         "0 second ago",
		"atm":         "am",
		"apr":         "april",
		"aug":         "august",
		"day":         "day",
		"dec":         "december",
		"feb":         "february",
		"gmt":         "gmt",
		"jan":         "january",
		"jul":         "july",
		"jun":         "june",
		"mar":         "march",
		"maj":         "may",
		"nov":         "november",
		"okt":         "october",
		"ptm":         "pm",
		"sep":         "september",
		"utc":         "utc",
		"am":          "am",
		"ve":          "friday",
		"lu":          "monday",
		"pm":          "pm",
		"sa":          "saturday",
		"di":          "sunday",
		"ja":          "thursday",
		"ma":          "tuesday",
		"me":          "wednesday",
		"'":           "",
		",":           "",
		";":           "",
		"@":           "",
		"[":           "",
		"]":           "",
		"|":           "",
		" ":           " ",
		"+":           "+",
		"-":           "-",
		".":           ".",
		"/":           "/",
		":":           ":",
		"z":           "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|septembro|this hour|this week|this year|yesterday|decembro|februaro|merkredo|novembro|tomorrow|vendredo|augusto|dimanco|januaro|oktobro|aprilo|minute|sabato|second|jaudo|julio|junio|lundo|mardo|marto|month|today|hour|majo|week|year|apr|atm|aug|day|dec|feb|gmt|jan|jul|jun|maj|mar|nov|now|okt|ptm|sep|utc|\+|\.|\[|\]|\||am|di|ja|lu|ma|me|pm|sa|ve| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

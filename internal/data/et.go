// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var et_Locale = merge(nil, LocaleData{
	Name:      "et",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"praegusel minutil": "0 minute ago",
		"praegusel tunnil":  "0 hour ago",
		"kaesolev nadal":    "0 week ago",
		"kaesolev aasta":    "0 year ago",
		"jargmine nadal":    "in 1 week",
		"jargmine aasta":    "in 1 year",
		"eelmine nadal":     "1 week ago",
		"eelmine aasta":     "1 year ago",
		"kaesolev kuu":      "0 month ago",
		"jargmine kuu":      "in 1 month",
		"eelmine kuu":       "1 month ago",
		"detsember":         "december",
		"esmaspaev":         "monday",
		"september":         "september",
		"neljapaev":         "thursday",
		"teisipaev":         "tuesday",
		"kolmapaev":         "wednesday",
		"veebruar":          "february",
		"november":          "november",
		"oktoober":          "october",
		"puhapaev":          "sunday",
		"jaanuar":           "january",
		"laupaev":           "saturday",
		"aprill":            "april",
		"august":            "august",
		"sekund":            "second",
		"veebr":             "february",
		"reede":             "friday",
		"homme":             "in 1 day",
		"juuli":             "july",
		"juuni":             "june",
		"marts":             "march",
		"minut":             "minute",
		"nadal":             "week",
		"aasta":             "year",
		"tana":              "0 day ago",
		"nuud":              "0 second ago",
		"eile":              "1 day ago",
		"paev":              "day",
		"dets":              "december",
		"tund":              "hour",
		"jaan":              "january",
		"sept":              "september",
		"apr":               "april",
		"aug":               "august",
		"gmt":               "gmt",
		"mai":               "may",
		"min":               "minute",
		"kuu":               "month",
		"nov":               "november",
		"okt":               "october",
		"sek":               "second",
		"utc":               "utc",
		"nad":               "week",
		"am":                "am",
		"pm":                "pm",
		"'":                 "",
		",":                 "",
		";":                 "",
		"@":                 "",
		"[":                 "",
		"]":                 "",
		"|":                 "",
		" ":                 " ",
		"+":                 "+",
		"-":                 "-",
		".":                 ".",
		"/":                 "/",
		":":                 ":",
		"r":                 "friday",
		"e":                 "monday",
		"l":                 "saturday",
		"s":                 "second",
		"p":                 "sunday",
		"n":                 "thursday",
		"t":                 "tuesday",
		"k":                 "wednesday",
		"a":                 "year",
		"z":                 "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) sekundi parast`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) minuti parast`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) nadala parast`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) sekundi eest`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) paeva parast`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) tunni parast`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) aasta parast`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) minuti eest`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) nadala eest`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) paeva eest`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) tunni eest`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) aasta eest`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) min parast`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) kuu parast`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) sek parast`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) nad parast`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) min eest`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) kuu eest`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) sek eest`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) nad eest`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) p parast`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) t parast`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) k parast`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) s parast`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) a parast`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) p eest`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) t eest`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) k eest`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) s eest`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) a eest`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ sekundi parast|\d+ minuti parast|\d+ nadala parast|\d+ aasta parast|\d+ paeva parast|\d+ sekundi eest|\d+ tunni parast|\d+ minuti eest|\d+ nadala eest|\d+ aasta eest|\d+ kuu parast|\d+ min parast|\d+ nad parast|\d+ paeva eest|\d+ sek parast|\d+ tunni eest|\d+ a parast|\d+ k parast|\d+ kuu eest|\d+ min eest|\d+ nad eest|\d+ p parast|\d+ s parast|\d+ sek eest|\d+ t parast|\d+ a eest|\d+ k eest|\d+ p eest|\d+ s eest|\d+ t eest)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ sekundi parast|\d+ minuti parast|\d+ nadala parast|\d+ aasta parast|\d+ paeva parast|\d+ sekundi eest|\d+ tunni parast|\d+ minuti eest|\d+ nadala eest|\d+ aasta eest|\d+ kuu parast|\d+ min parast|\d+ nad parast|\d+ paeva eest|\d+ sek parast|\d+ tunni eest|\d+ a parast|\d+ k parast|\d+ kuu eest|\d+ min eest|\d+ nad eest|\d+ p parast|\d+ s parast|\d+ sek eest|\d+ t parast|\d+ a eest|\d+ k eest|\d+ p eest|\d+ s eest|\d+ t eest)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(praegusel minutil|praegusel tunnil|jargmine aasta|jargmine nadal|kaesolev aasta|kaesolev nadal|eelmine aasta|eelmine nadal|jargmine kuu|kaesolev kuu|eelmine kuu|detsember|esmaspaev|kolmapaev|neljapaev|september|teisipaev|november|oktoober|puhapaev|veebruar|jaanuar|laupaev|aprill|august|sekund|aasta|homme|juuli|juuni|marts|minut|nadal|reede|veebr|dets|eile|jaan|nuud|paev|sept|tana|tund|apr|aug|gmt|kuu|mai|min|nad|nov|okt|sek|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|a|e|k|l|n|p|r|s|t|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

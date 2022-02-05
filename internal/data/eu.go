// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var eu_Locale = merge(nil, LocaleData{
	Name:      "eu",
	DateOrder: "YMD",
	Charset:   []rune(`bcdeghiklnorstuxz`),
	Translations: map[string]string{
		"asteazkena": "wednesday",
		"astelehena": "monday",
		"asteartea":  "tuesday",
		"hilabetea":  "month",
		"larunbata":  "saturday",
		"urtarrila":  "january",
		"osteguna":   "thursday",
		"ostirala":   "friday",
		"segundoa":   "second",
		"abendua":    "december",
		"abuztua":    "august",
		"apirila":    "april",
		"igandea":    "sunday",
		"maiatza":    "may",
		"martxoa":    "march",
		"minutua":    "minute",
		"otsaila":    "february",
		"uztaila":    "july",
		"azaroa":     "november",
		"ekaina":     "june",
		"iraila":     "september",
		"astea":      "week",
		"eguna":      "day",
		"ordua":      "hour",
		"urria":      "october",
		"urtea":      "year",
		"abe":        "december",
		"abu":        "august",
		"api":        "april",
		"ast":        "week",
		"aza":        "november",
		"eka":        "june",
		"gmt":        "gmt",
		"hil":        "month",
		"ira":        "september",
		"mai":        "may",
		"mar":        "march",
		"min":        "minute",
		"ots":        "february",
		"urr":        "october",
		"urt":        "january",
		"utc":        "utc",
		"uzt":        "july",
		"al":         "monday",
		"am":         "am",
		"ar":         "tuesday",
		"az":         "wednesday",
		"eg":         "day",
		"ig":         "sunday",
		"lr":         "saturday",
		"og":         "thursday",
		"or":         "friday",
		"pm":         "pm",
		" ":          " ",
		"'":          "",
		"+":          "+",
		",":          "",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"h":          "hour",
		"s":          "second",
		"z":          "z",
		"|":          "",
	},
	RelativeType: map[string]string{
		"hurrengo hilabetea": "in 1 month",
		"aurreko hilabetea":  "1 month ago",
		"hurrengo astea":     "in 1 week",
		"hurrengo urtea":     "in 1 year",
		"minutu honetan":     "0 minute ago",
		"aurreko astea":      "1 week ago",
		"aurreko urtea":      "1 year ago",
		"hilabete hau":       "0 month ago",
		"ordu honetan":       "0 hour ago",
		"aste hau":           "0 week ago",
		"aurten":             "0 year ago",
		"bihar":              "in 1 day",
		"orain":              "0 second ago",
		"atzo":               "1 day ago",
		"gaur":               "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) hilabete barru`), "in $1 month"},
		{regexp.MustCompile(`(?i)duela (\d+) hilabete`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) segundo barru`), "in $1 second"},
		{regexp.MustCompile(`(?i)duela (\d+) segundo`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) minutu barru`), "in $1 minute"},
		{regexp.MustCompile(`(?i)duela (\d+) minutu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) aste barru`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) egun barru`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ordu barru`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) urte barru`), "in $1 year"},
		{regexp.MustCompile(`(?i)duela (\d+) aste`), "$1 week ago"},
		{regexp.MustCompile(`(?i)duela (\d+) egun`), "$1 day ago"},
		{regexp.MustCompile(`(?i)duela (\d+) ordu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)duela (\d+) urte`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ hilabete barru|duela \d+ hilabete|\d+ segundo barru|duela \d+ segundo|\d+ minutu barru|duela \d+ minutu|\d+ aste barru|\d+ egun barru|\d+ ordu barru|\d+ urte barru|duela \d+ aste|duela \d+ egun|duela \d+ ordu|duela \d+ urte)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ hilabete barru|duela \d+ hilabete|\d+ segundo barru|duela \d+ segundo|\d+ minutu barru|duela \d+ minutu|\d+ aste barru|\d+ egun barru|\d+ ordu barru|\d+ urte barru|duela \d+ aste|duela \d+ egun|duela \d+ ordu|duela \d+ urte)$`),
	KnownWords:      []string{"hurrengo hilabetea", "aurreko hilabetea", "hurrengo astea", "hurrengo urtea", "minutu honetan", "aurreko astea", "aurreko urtea", "hilabete hau", "ordu honetan", "asteazkena", "astelehena", "asteartea", "hilabetea", "larunbata", "urtarrila", "aste hau", "osteguna", "ostirala", "segundoa", "abendua", "abuztua", "apirila", "igandea", "maiatza", "martxoa", "minutua", "otsaila", "uztaila", "aurten", "azaroa", "ekaina", "iraila", "astea", "bihar", "eguna", "orain", "ordua", "urria", "urtea", "atzo", "gaur", "abe", "abu", "api", "ast", "aza", "eka", "gmt", "hil", "ira", "mai", "mar", "min", "ots", "urr", "urt", "utc", "uzt", "al", "am", "ar", "az", "eg", "ig", "lr", "og", "or", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "h", "s", "z", "|"},
})

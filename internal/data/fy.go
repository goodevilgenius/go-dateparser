// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fy_Locale = merge(nil, LocaleData{
	Name:      "fy",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"folgjende moanne": "in 1 month",
		"foarige moanne":   "1 month ago",
		"folgjende wike":   "in 1 week",
		"folgjend jier":    "in 1 year",
		"dizze moanne":     "0 month ago",
		"foarige wike":     "1 week ago",
		"foarich jier":     "1 year ago",
		"this minute":      "0 minute ago",
		"dizze wike":       "0 week ago",
		"febrewaris":       "february",
		"jannewaris":       "january",
		"tongersdei":       "thursday",
		"this hour":        "0 hour ago",
		"septimber":        "september",
		"dit jier":         "0 year ago",
		"gisteren":         "1 day ago",
		"augustus":         "august",
		"desimber":         "december",
		"novimber":         "november",
		"woansdei":         "wednesday",
		"vandaag":          "0 day ago",
		"moandei":          "monday",
		"oktober":          "october",
		"sekonde":          "second",
		"tiisdei":          "tuesday",
		"morgen":           "in 1 day",
		"moanne":           "month",
		"april":            "april",
		"freed":            "friday",
		"maart":            "march",
		"maaie":            "may",
		"minut":            "minute",
		"sneon":            "saturday",
		"snein":            "sunday",
		"oere":             "hour",
		"july":             "july",
		"juny":             "june",
		"wike":             "week",
		"jier":             "year",
		"apr":              "april",
		"aug":              "august",
		"dei":              "day",
		"des":              "december",
		"feb":              "february",
		"gmt":              "gmt",
		"jan":              "january",
		"jul":              "july",
		"jun":              "june",
		"mrt":              "march",
		"mai":              "may",
		"nov":              "november",
		"okt":              "october",
		"sep":              "september",
		"utc":              "utc",
		"nu":               "0 second ago",
		"am":               "am",
		"fr":               "friday",
		"mo":               "monday",
		"pm":               "pm",
		"so":               "saturday",
		"si":               "sunday",
		"to":               "thursday",
		"ti":               "tuesday",
		"wo":               "wednesday",
		"'":                "",
		",":                "",
		";":                "",
		"@":                "",
		"[":                "",
		"]":                "",
		"|":                "",
		" ":                " ",
		"+":                "+",
		"-":                "-",
		".":                ".",
		"/":                "/",
		":":                ":",
		"z":                "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) sekonden lyn`), "$1 second ago"},
		{regexp.MustCompile(`(?i)oer (\d+) sekonden`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) minuten lyn`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) moannen lyn`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) sekonde lyn`), "$1 second ago"},
		{regexp.MustCompile(`(?i)oer (\d+) minuten`), "in $1 minute"},
		{regexp.MustCompile(`(?i)oer (\d+) moannen`), "in $1 month"},
		{regexp.MustCompile(`(?i)oer (\d+) sekonde`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) moanne lyn`), "$1 month ago"},
		{regexp.MustCompile(`(?i)oer (\d+) moanne`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) deien lyn`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) minut lyn`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) wiken lyn`), "$1 week ago"},
		{regexp.MustCompile(`(?i)oer (\d+) deien`), "in $1 day"},
		{regexp.MustCompile(`(?i)oer (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)oer (\d+) wiken`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) oere lyn`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) wike lyn`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) jier lyn`), "$1 year ago"},
		{regexp.MustCompile(`(?i)oer (\d+) oere`), "in $1 hour"},
		{regexp.MustCompile(`(?i)oer (\d+) wike`), "in $1 week"},
		{regexp.MustCompile(`(?i)oer (\d+) jier`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) dei lyn`), "$1 day ago"},
		{regexp.MustCompile(`(?i)oer (\d+) dei`), "in $1 day"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ sekonden lyn|oer \d+ sekonden|\d+ minuten lyn|\d+ moannen lyn|\d+ sekonde lyn|oer \d+ minuten|oer \d+ moannen|oer \d+ sekonde|\d+ moanne lyn|oer \d+ moanne|\d+ deien lyn|\d+ minut lyn|\d+ wiken lyn|oer \d+ deien|oer \d+ minut|oer \d+ wiken|\d+ jier lyn|\d+ oere lyn|\d+ wike lyn|oer \d+ jier|oer \d+ oere|oer \d+ wike|\d+ dei lyn|oer \d+ dei)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ sekonden lyn|oer \d+ sekonden|\d+ minuten lyn|\d+ moannen lyn|\d+ sekonde lyn|oer \d+ minuten|oer \d+ moannen|oer \d+ sekonde|\d+ moanne lyn|oer \d+ moanne|\d+ deien lyn|\d+ minut lyn|\d+ wiken lyn|oer \d+ deien|oer \d+ minut|oer \d+ wiken|\d+ jier lyn|\d+ oere lyn|\d+ wike lyn|oer \d+ jier|oer \d+ oere|oer \d+ wike|\d+ dei lyn|oer \d+ dei)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(folgjende moanne|foarige moanne|folgjende wike|folgjend jier|dizze moanne|foarich jier|foarige wike|this minute|dizze wike|febrewaris|jannewaris|tongersdei|septimber|this hour|augustus|desimber|dit jier|gisteren|novimber|woansdei|moandei|oktober|sekonde|tiisdei|vandaag|moanne|morgen|april|freed|maaie|maart|minut|snein|sneon|jier|july|juny|oere|wike|apr|aug|dei|des|feb|gmt|jan|jul|jun|mai|mrt|nov|okt|sep|utc|\+|\.|\[|\]|\||am|fr|mo|nu|pm|si|so|ti|to|wo| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

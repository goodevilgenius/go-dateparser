// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var teo_Locale = merge(nil, LocaleData{
	Name:      "teo",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"omodok'king'ol": "june",
		"this minute":    "0 minute ago",
		"nakaebarasa":    "monday",
		"osokosokoma":    "september",
		"this month":     "0 month ago",
		"last month":     "1 month ago",
		"next month":     "in 1 month",
		"nakasabiti":     "saturday",
		"nakaung'on":     "thursday",
		"this hour":      "0 hour ago",
		"this week":      "0 week ago",
		"this year":      "0 year ago",
		"last week":      "1 week ago",
		"last year":      "1 year ago",
		"taparachu":      "am",
		"next week":      "in 1 week",
		"next year":      "in 1 year",
		"nakaejuma":      "sunday",
		"odung'el":       "april",
		"nakakany":       "friday",
		"isekonde":       "second",
		"okwamg'":        "march",
		"idakika":        "minute",
		"nakaare":        "tuesday",
		"nakauni":        "wednesday",
		"opedel":         "august",
		"aparan":         "day",
		"omaruk":         "may",
		"olabor":         "november",
		"otibar":         "october",
		"ebongi":         "pm",
		"orara":          "january",
		"ojola":          "july",
		"ewiki":          "week",
		"lolo":           "0 day ago",
		"jaan":           "1 day ago",
		"opoo":           "december",
		"omuk":           "february",
		"esaa":           "hour",
		"elap":           "month",
		"ekan":           "year",
		"now":            "0 second ago",
		"dun":            "april",
		"ped":            "august",
		"poo":            "december",
		"muk":            "february",
		"kan":            "friday",
		"gmt":            "gmt",
		"moi":            "in 1 day",
		"rar":            "january",
		"jol":            "july",
		"mod":            "june",
		"kwa":            "march",
		"mar":            "may",
		"bar":            "monday",
		"lab":            "november",
		"tib":            "october",
		"sab":            "saturday",
		"sok":            "september",
		"jum":            "sunday",
		"ung":            "thursday",
		"aar":            "tuesday",
		"utc":            "utc",
		"uni":            "wednesday",
		"am":             "am",
		"pm":             "pm",
		"'":              "",
		",":              "",
		";":              "",
		"@":              "",
		"[":              "",
		"]":              "",
		"|":              "",
		" ":              " ",
		"+":              "+",
		"-":              "-",
		".":              ".",
		"/":              "/",
		":":              ":",
		"z":              "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(omodok'king'ol|nakaebarasa|osokosokoma|this minute|last month|nakasabiti|nakaung'on|next month|this month|last week|last year|nakaejuma|next week|next year|taparachu|this hour|this week|this year|isekonde|nakakany|odung'el|idakika|nakaare|nakauni|okwamg'|aparan|ebongi|olabor|omaruk|opedel|otibar|ewiki|ojola|orara|ekan|elap|esaa|jaan|lolo|omuk|opoo|aar|bar|dun|gmt|jol|jum|kan|kwa|lab|mar|mod|moi|muk|now|ped|poo|rar|sab|sok|tib|ung|uni|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var teo_KE_Locale = merge(&teo_Locale, LocaleData{
	Name:      "teo-KE",
	DateOrder: "DMY",
})

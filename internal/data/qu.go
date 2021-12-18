// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var qu_Locale = merge(nil, LocaleData{
	Name:      "qu",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"qhapaq sitwa": "august",
		"pauqar waray": "march",
		"this minute":  "0 minute ago",
		"kapaq raymi":  "december",
		"hatun puquy":  "february",
		"qulla puquy":  "january",
		"this month":   "0 month ago",
		"last month":   "1 month ago",
		"next month":   "in 1 month",
		"anta sitwa":   "july",
		"inti raymi":   "june",
		"this hour":    "0 hour ago",
		"this week":    "0 week ago",
		"this year":    "0 year ago",
		"yesterday":    "1 day ago",
		"last week":    "1 week ago",
		"last year":    "1 year ago",
		"next week":    "in 1 week",
		"next year":    "in 1 year",
		"ayamarq'a":    "november",
		"uma raymi":    "september",
		"miercoles":    "wednesday",
		"tomorrow":     "in 1 day",
		"kantaray":     "october",
		"viernes":      "friday",
		"aymuray":      "may",
		"domingo":      "sunday",
		"ayriwa":       "april",
		"minute":       "minute",
		"sabado":       "saturday",
		"second":       "second",
		"jueves":       "thursday",
		"martes":       "tuesday",
		"today":        "0 day ago",
		"lunes":        "monday",
		"month":        "month",
		"hour":         "hour",
		"week":         "week",
		"year":         "year",
		"now":          "0 second ago",
		"ayr":          "april",
		"qha":          "august",
		"day":          "day",
		"kap":          "december",
		"hat":          "february",
		"vie":          "friday",
		"gmt":          "gmt",
		"qul":          "january",
		"ant":          "july",
		"int":          "june",
		"pau":          "march",
		"aym":          "may",
		"lun":          "monday",
		"aya":          "november",
		"kan":          "october",
		"sab":          "saturday",
		"uma":          "september",
		"dom":          "sunday",
		"jue":          "thursday",
		"mar":          "tuesday",
		"utc":          "utc",
		"mie":          "wednesday",
		"am":           "am",
		"pm":           "pm",
		"'":            "",
		",":            "",
		";":            "",
		"@":            "",
		"[":            "",
		"]":            "",
		"|":            "",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"z":            "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(pauqar waray|qhapaq sitwa|hatun puquy|kapaq raymi|qulla puquy|this minute|anta sitwa|inti raymi|last month|next month|this month|ayamarq'a|last week|last year|miercoles|next week|next year|this hour|this week|this year|uma raymi|yesterday|kantaray|tomorrow|aymuray|domingo|viernes|ayriwa|jueves|martes|minute|sabado|second|lunes|month|today|hour|week|year|ant|aya|aym|ayr|day|dom|gmt|hat|int|jue|kan|kap|lun|mar|mie|now|pau|qha|qul|sab|uma|utc|vie|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var qu_BO_Locale = merge(&qu_Locale, LocaleData{
	Name:      "qu-BO",
	DateOrder: "DMY",
})

var qu_EC_Locale = merge(&qu_Locale, LocaleData{
	Name:      "qu-EC",
	DateOrder: "DMY",
})

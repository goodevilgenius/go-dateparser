// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mt_Locale = merge(nil, LocaleData{
	Name:      "mt",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"il-gimgħa li għaddiet": "1 week ago",
		"is-sena li għaddiet":   "1 year ago",
		"ix-xahar li għadda":    "1 month ago",
		"ix-xahar id-dieħel":    "in 1 month",
		"il-gimgħa d-dieħla":    "in 1 week",
		"is-sena d-dieħla":      "in 1 year",
		"din il-gimgħa":         "0 week ago",
		"dan ix-xahar":          "0 month ago",
		"this minute":           "0 minute ago",
		"din is-sena":           "0 year ago",
		"this hour":             "0 hour ago",
		"il-gimgħa":             "friday",
		"settembru":             "september",
		"it-tlieta":             "tuesday",
		"ilbieraħ":              "1 day ago",
		"dicembru":              "december",
		"it-tnejn":              "monday",
		"novembru":              "november",
		"il-ħamis":              "thursday",
		"l-erbgħa":              "wednesday",
		"awwissu":               "august",
		"ottubru":               "october",
		"is-sibt":               "saturday",
		"sekonda":               "second",
		"il-ħadd":               "sunday",
		"siegħa":                "hour",
		"jannar":                "january",
		"minuta":                "minute",
		"gimgħa":                "week",
		"illum":                 "0 day ago",
		"april":                 "april",
		"għada":                 "in 1 day",
		"lulju":                 "july",
		"gunju":                 "june",
		"marzu":                 "march",
		"mejju":                 "may",
		"xahar":                 "month",
		"frar":                  "february",
		"sena":                  "year",
		"now":                   "0 second ago",
		"apr":                   "april",
		"aww":                   "august",
		"jum":                   "day",
		"dic":                   "december",
		"fra":                   "february",
		"gim":                   "friday",
		"gmt":                   "gmt",
		"jan":                   "january",
		"lul":                   "july",
		"gun":                   "june",
		"mar":                   "march",
		"mej":                   "may",
		"tne":                   "monday",
		"nov":                   "november",
		"ott":                   "october",
		"sib":                   "saturday",
		"set":                   "september",
		"ħad":                   "sunday",
		"ħam":                   "thursday",
		"tli":                   "tuesday",
		"utc":                   "utc",
		"erb":                   "wednesday",
		"am":                    "am",
		"pm":                    "pm",
		"'":                     "",
		",":                     "",
		";":                     "",
		"@":                     "",
		"[":                     "",
		"]":                     "",
		"|":                     "",
		" ":                     " ",
		"+":                     "+",
		"-":                     "-",
		".":                     ".",
		"/":                     "/",
		":":                     ":",
		"z":                     "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) sena ilu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) snin ilu`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ sena ilu|\d+ snin ilu)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ sena ilu|\d+ snin ilu)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(il-gimgħa li għaddiet|is-sena li għaddiet|il-gimgħa d-dieħla|ix-xahar id-dieħel|ix-xahar li għadda|is-sena d-dieħla|din il-gimgħa|dan ix-xahar|din is-sena|this minute|il-gimgħa|it-tlieta|settembru|this hour|dicembru|il-ħamis|ilbieraħ|it-tnejn|l-erbgħa|novembru|awwissu|il-ħadd|is-sibt|ottubru|sekonda|gimgħa|jannar|minuta|siegħa|april|gunju|għada|illum|lulju|marzu|mejju|xahar|frar|sena|apr|aww|dic|erb|fra|gim|gmt|gun|jan|jum|lul|mar|mej|nov|now|ott|set|sib|tli|tne|utc|ħad|ħam|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

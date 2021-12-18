// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var km_Locale = merge(nil, LocaleData{
	Name:      "km",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"សបដាហ​ករោយ": "in 1 week",
		"សបដាហ​នេះ":  "0 week ago",
		"សបដាហ​មន":   "1 week ago",
		"ថងៃ​សអែក":   "in 1 day",
		"ឆនា​ករោយ":   "in 1 year",
		"ថងៃ​នេះ":    "0 day ago",
		"ឆនា​នេះ":    "0 year ago",
		"ថងៃសអែក":    "in 1 day",
		"ខែ​ករោយ":    "in 1 month",
		"មោងនេះ":     "0 hour ago",
		"នាទនេះ":     "0 minute ago",
		"ខែ​នេះ":     "0 month ago",
		"ឆនា​មន":     "1 year ago",
		"ពរហសបត":     "thursday",
		"មសលមញ":      "1 day ago",
		"ខែ​មន":      "1 month ago",
		"កកកដា":      "july",
		"វចឆកា":      "november",
		"អាទតយ":      "sunday",
		"អងគារ":      "tuesday",
		"សបដាហ":      "week",
		"មេសា":       "april",
		"កមភៈ":       "february",
		"មករា":       "january",
		"មថនា":       "june",
		"ឧសភា":       "may",
		"វនាទ":       "second",
		"កញញា":       "september",
		"ឥឡវ":        "0 second ago",
		"សហា":        "august",
		"ថងៃ":        "day",
		"សករ":        "friday",
		"gmt":        "gmt",
		"មោង":        "hour",
		"មនា":        "march",
		"នាទ":        "minute",
		"ចនទ":        "monday",
		"តលា":        "october",
		"សៅរ":        "saturday",
		"utc":        "utc",
		"ឆនា":        "year",
		"am":         "am",
		"ធន":         "december",
		"ខែ":         "month",
		"pm":         "pm",
		"ពធ":         "wednesday",
		"'":          "",
		",":          "",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"|":          "",
		" ":          " ",
		"+":          "+",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		"z":          "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)កនង​រយៈ​ពេល (\d+) មោង`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) សបដាហ​មន`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) សបដាហទៀត`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ថងៃ​​មន`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) នាទ​​មន`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) វនាទ​មន`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) វនាទទៀត`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) ថងៃ​មន`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) មោង​មន`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) នាទ​មន`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ឆនា​មន`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) ថងៃទៀត`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) មោងទៀត`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) នាទទៀត`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ឆនាទៀត`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ខែទៀត`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) ខែមន`), "$1 month ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(កនង​រយៈ​ពេល \d+ មោង|\d+ សបដាហទៀត|\d+ សបដាហ​មន|\d+ ថងៃ​​មន|\d+ នាទ​​មន|\d+ វនាទទៀត|\d+ វនាទ​មន|\d+ ឆនាទៀត|\d+ ឆនា​មន|\d+ ថងៃទៀត|\d+ ថងៃ​មន|\d+ នាទទៀត|\d+ នាទ​មន|\d+ មោងទៀត|\d+ មោង​មន|\d+ ខែទៀត|\d+ ខែមន)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(កនង​រយៈ​ពេល \d+ មោង|\d+ សបដាហទៀត|\d+ សបដាហ​មន|\d+ ថងៃ​​មន|\d+ នាទ​​មន|\d+ វនាទទៀត|\d+ វនាទ​មន|\d+ ឆនាទៀត|\d+ ឆនា​មន|\d+ ថងៃទៀត|\d+ ថងៃ​មន|\d+ នាទទៀត|\d+ នាទ​មន|\d+ មោងទៀត|\d+ មោង​មន|\d+ ខែទៀត|\d+ ខែមន)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(សបដាហ​ករោយ|សបដាហ​នេះ|ឆនា​ករោយ|ថងៃ​សអែក|សបដាហ​មន|ខែ​ករោយ|ឆនា​នេះ|ថងៃសអែក|ថងៃ​នេះ|ខែ​នេះ|ឆនា​មន|នាទនេះ|ពរហសបត|មោងនេះ|កកកដា|ខែ​មន|មសលមញ|វចឆកា|សបដាហ|អងគារ|អាទតយ|កញញា|កមភៈ|មករា|មថនា|មេសា|វនាទ|ឧសភា|gmt|utc|ចនទ|ឆនា|តលា|ថងៃ|នាទ|មនា|មោង|សករ|សហា|សៅរ|ឥឡវ|\+|\.|\[|\]|\||am|pm|ខែ|ធន|ពធ| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

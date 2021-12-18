// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var or_Locale = merge(nil, LocaleData{
	Name:      "or",
	DateOrder: "DMY",
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
		"tomorrow":    "in 1 day",
		"ସେପଟେମବର":    "september",
		"ମଙଗଳବାର":     "tuesday",
		"ଡସେମବର":      "december",
		"ଫେବଆରୀ":      "february",
		"ଶକରବାର":      "friday",
		"ଜାନଆରୀ":      "january",
		"minute":      "minute",
		"ସୋମବାର":      "monday",
		"ନଭେମବର":      "november",
		"ଅକଟୋବର":      "october",
		"second":      "second",
		"today":       "0 day ago",
		"ଅପରେଲ":       "april",
		"ମାରଚଚ":       "march",
		"month":       "month",
		"ଶନବାର":       "saturday",
		"ରବବାର":       "sunday",
		"ଗରବାର":       "thursday",
		"ବଧବାର":       "wednesday",
		"ଅଗଷଟ":        "august",
		"hour":        "hour",
		"ଜଲାଇ":        "july",
		"ମଙଗଳ":        "tuesday",
		"week":        "week",
		"year":        "year",
		"now":         "0 second ago",
		"day":         "day",
		"ଶକର":         "friday",
		"gmt":         "gmt",
		"ସୋମ":         "monday",
		"utc":         "utc",
		"am":          "am",
		"ଜନ":          "june",
		"ମଇ":          "may",
		"pm":          "pm",
		"ଶନ":          "saturday",
		"ରବ":          "sunday",
		"ଗର":          "thursday",
		"ବଧ":          "wednesday",
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
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|yesterday|tomorrow|ସେପଟେମବର|ମଙଗଳବାର|minute|second|ଅକଟୋବର|ଜାନଆରୀ|ଡସେମବର|ନଭେମବର|ଫେବଆରୀ|ଶକରବାର|ସୋମବାର|month|today|ଅପରେଲ|ଗରବାର|ବଧବାର|ମାରଚଚ|ରବବାର|ଶନବାର|hour|week|year|ଅଗଷଟ|ଜଲାଇ|ମଙଗଳ|day|gmt|now|utc|ଶକର|ସୋମ|\+|\.|\[|\]|\||am|pm|ଗର|ଜନ|ବଧ|ମଇ|ରବ|ଶନ| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

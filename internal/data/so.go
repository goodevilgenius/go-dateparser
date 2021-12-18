// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var so_Locale = merge(nil, LocaleData{
	Name:      "so",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"bisha laba iyo tobnaad": "december",
		"bisha kow iyo tobnaad":  "november",
		"bisha sideedaad":        "august",
		"bisha saddexaad":        "march",
		"bisha sagaalaad":        "september",
		"bisha todobaad":         "july",
		"bisha koobaad":          "january",
		"bisha shanaad":          "may",
		"bisha tobnaad":          "october",
		"bisha afraad":           "april",
		"bisha labaad":           "february",
		"bisha lixaad":           "june",
		"this minute":            "0 minute ago",
		"this month":             "0 month ago",
		"last month":             "1 month ago",
		"next month":             "in 1 month",
		"this hour":              "0 hour ago",
		"this week":              "0 week ago",
		"this year":              "0 year ago",
		"last week":              "1 week ago",
		"last year":              "1 year ago",
		"next week":              "in 1 week",
		"next year":              "in 1 year",
		"khamiis":                "thursday",
		"talaado":                "tuesday",
		"maanta":                 "0 day ago",
		"shalay":                 "1 day ago",
		"minute":                 "minute",
		"isniin":                 "monday",
		"second":                 "second",
		"arbaco":                 "wednesday",
		"jimco":                  "friday",
		"berri":                  "in 1 day",
		"month":                  "month",
		"sabti":                  "saturday",
		"hour":                   "hour",
		"axad":                   "sunday",
		"week":                   "week",
		"year":                   "year",
		"now":                    "0 second ago",
		"afr":                    "april",
		"sid":                    "august",
		"day":                    "day",
		"lit":                    "december",
		"lab":                    "february",
		"jim":                    "friday",
		"gmt":                    "gmt",
		"kob":                    "january",
		"tod":                    "july",
		"lix":                    "june",
		"sad":                    "march",
		"sha":                    "may",
		"isn":                    "monday",
		"kit":                    "november",
		"tob":                    "october",
		"sab":                    "saturday",
		"sag":                    "september",
		"axd":                    "sunday",
		"kha":                    "thursday",
		"tal":                    "tuesday",
		"utc":                    "utc",
		"arb":                    "wednesday",
		"am":                     "am",
		"sn":                     "am",
		"gn":                     "pm",
		"pm":                     "pm",
		"'":                      "",
		",":                      "",
		";":                      "",
		"@":                      "",
		"[":                      "",
		"]":                      "",
		"|":                      "",
		" ":                      " ",
		"+":                      "+",
		"-":                      "-",
		".":                      ".",
		"/":                      "/",
		":":                      ":",
		"z":                      "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(bisha laba iyo tobnaad|bisha kow iyo tobnaad|bisha saddexaad|bisha sagaalaad|bisha sideedaad|bisha todobaad|bisha koobaad|bisha shanaad|bisha tobnaad|bisha afraad|bisha labaad|bisha lixaad|this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|khamiis|talaado|arbaco|isniin|maanta|minute|second|shalay|berri|jimco|month|sabti|axad|hour|week|year|afr|arb|axd|day|gmt|isn|jim|kha|kit|kob|lab|lit|lix|now|sab|sad|sag|sha|sid|tal|tob|tod|utc|\+|\.|\[|\]|\||am|gn|pm|sn| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var so_DJ_Locale = merge(&so_Locale, LocaleData{
	Name:      "so-DJ",
	DateOrder: "DMY",
})

var so_ET_Locale = merge(&so_Locale, LocaleData{
	Name:      "so-ET",
	DateOrder: "DMY",
})

var so_KE_Locale = merge(&so_Locale, LocaleData{
	Name:      "so-KE",
	DateOrder: "DMY",
})

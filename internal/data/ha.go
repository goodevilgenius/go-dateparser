// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ha_Locale = merge(nil, LocaleData{
	Name:      "ha",
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
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"faburairu":   "february",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"afirilu":     "april",
		"disamba":     "december",
		"jumma'a":     "friday",
		"janairu":     "january",
		"litinin":     "monday",
		"nuwamba":     "november",
		"satumba":     "september",
		"alhamis":     "thursday",
		"shekara":     "year",
		"agusta":      "august",
		"oktoba":      "october",
		"asabar":      "saturday",
		"daƙiƙa":      "second",
		"lahadi":      "sunday",
		"talata":      "tuesday",
		"laraba":      "wednesday",
		"kwana":       "day",
		"maris":       "march",
		"minti":       "minute",
		"jiya":        "1 day ago",
		"gobe":        "in 1 day",
		"yuli":        "july",
		"yuni":        "june",
		"mayu":        "may",
		"wata":        "month",
		"mako":        "week",
		"yau":         "0 day ago",
		"now":         "0 second ago",
		"afi":         "april",
		"agu":         "august",
		"dis":         "december",
		"fab":         "february",
		"jum":         "friday",
		"gmt":         "gmt",
		"awa":         "hour",
		"jan":         "january",
		"yul":         "july",
		"yun":         "june",
		"mar":         "march",
		"may":         "may",
		"lit":         "monday",
		"nuw":         "november",
		"okt":         "october",
		"asa":         "saturday",
		"sat":         "september",
		"lah":         "sunday",
		"alh":         "thursday",
		"tal":         "tuesday",
		"utc":         "utc",
		"lar":         "wednesday",
		"am":          "am",
		"pm":          "pm",
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
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|faburairu|last week|last year|next week|next year|this hour|this week|this year|afirilu|alhamis|disamba|janairu|jumma'a|litinin|nuwamba|satumba|shekara|agusta|asabar|daƙiƙa|lahadi|laraba|oktoba|talata|kwana|maris|minti|gobe|jiya|mako|mayu|wata|yuli|yuni|afi|agu|alh|asa|awa|dis|fab|gmt|jan|jum|lah|lar|lit|mar|may|now|nuw|okt|sat|tal|utc|yau|yul|yun|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ha_GH_Locale = merge(&ha_Locale, LocaleData{
	Name:      "ha-GH",
	DateOrder: "DMY",
})

var ha_NE_Locale = merge(&ha_Locale, LocaleData{
	Name:      "ha-NE",
	DateOrder: "DMY",
})

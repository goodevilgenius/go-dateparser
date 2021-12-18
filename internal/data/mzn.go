// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mzn_Locale = merge(nil, LocaleData{
	Name:      "mzn",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"this minute": "0 minute ago",
		"this hour":   "0 hour ago",
		"قبلی هفته":   "1 week ago",
		"بعدی هفته":   "in 1 week",
		"این هفته":    "0 week ago",
		"سال دیگه":    "in 1 year",
		"این ماه":     "0 month ago",
		"ماه قبل":     "1 month ago",
		"ماه بعد":     "in 1 month",
		"سپتامبر":     "september",
		"پارسال":      "1 year ago",
		"دسامبر":      "december",
		"ژانویه":      "january",
		"نوامبر":      "november",
		"امروز":       "0 day ago",
		"امسال":       "0 year ago",
		"دیروز":       "1 day ago",
		"اوریل":       "april",
		"فوریه":       "february",
		"ژويیه":       "july",
		"دقیقه":       "minute",
		"اکتبر":       "october",
		"ثانیه":       "second",
		"ساعت":        "hour",
		"فردا":        "in 1 day",
		"ژوين":        "june",
		"مارس":        "march",
		"هفته":        "week",
		"now":         "0 second ago",
		"اوت":         "august",
		"روز":         "day",
		"fri":         "friday",
		"gmt":         "gmt",
		"دقه":         "minute",
		"mon":         "monday",
		"ماه":         "month",
		"sat":         "saturday",
		"sun":         "sunday",
		"thu":         "thursday",
		"tue":         "tuesday",
		"utc":         "utc",
		"wed":         "wednesday",
		"سال":         "year",
		"am":          "am",
		"مه":          "may",
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
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) ثانیه پیش`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) دقیقه دله`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ثانیه دله`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) ساعت پیش`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) هفته پیش`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ساعت دله`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) هفته دله`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) روز پیش`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) دقه پیش`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ماه پیش`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) سال پیش`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) روز دله`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) دقه دله`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ماه دله`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) سال دله`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ ثانیه دله|\d+ ثانیه پیش|\d+ دقیقه دله|\d+ ساعت دله|\d+ ساعت پیش|\d+ هفته دله|\d+ هفته پیش|\d+ دقه دله|\d+ دقه پیش|\d+ روز دله|\d+ روز پیش|\d+ سال دله|\d+ سال پیش|\d+ ماه دله|\d+ ماه پیش)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ ثانیه دله|\d+ ثانیه پیش|\d+ دقیقه دله|\d+ ساعت دله|\d+ ساعت پیش|\d+ هفته دله|\d+ هفته پیش|\d+ دقه دله|\d+ دقه پیش|\d+ روز دله|\d+ روز پیش|\d+ سال دله|\d+ سال پیش|\d+ ماه دله|\d+ ماه پیش)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|this hour|بعدی هفته|قبلی هفته|این هفته|سال دیگه|این ماه|سپتامبر|ماه بعد|ماه قبل|دسامبر|نوامبر|پارسال|ژانویه|امروز|امسال|اوریل|اکتبر|ثانیه|دقیقه|دیروز|فوریه|ژويیه|ساعت|فردا|مارس|هفته|ژوين|fri|gmt|mon|now|sat|sun|thu|tue|utc|wed|اوت|دقه|روز|سال|ماه|\+|\.|\[|\]|\||am|pm|مه| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

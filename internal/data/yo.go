// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var yo_Locale = merge(nil, LocaleData{
	Name:      "yo",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"this minute": "0 minute ago",
		"ojo abameta": "saturday",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"iseju aaya":  "second",
		"ojo isegun":  "tuesday",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"osu erele":   "february",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"osu agemo":   "july",
		"osu okudu":   "june",
		"osu erena":   "march",
		"osu ebibi":   "may",
		"osu owara":   "october",
		"osu owewe":   "september",
		"osu igbe":    "april",
		"osu ogun":    "august",
		"osu sere":    "january",
		"osu belu":    "november",
		"ojo aiku":    "sunday",
		"osu ope":     "december",
		"ojo eti":     "friday",
		"ojo aje":     "monday",
		"abameta":     "saturday",
		"wakati":      "hour",
		"isegun":      "tuesday",
		"erele":       "february",
		"agemo":       "july",
		"okudu":       "june",
		"erena":       "march",
		"ebibi":       "may",
		"iseju":       "minute",
		"owara":       "october",
		"owewe":       "september",
		"ojobo":       "thursday",
		"ojoru":       "wednesday",
		"aaro":        "am",
		"igbe":        "april",
		"ogun":        "august",
		"sere":        "january",
		"belu":        "november",
		"osan":        "pm",
		"aiku":        "sunday",
		"odun":        "year",
		"oni":         "0 day ago",
		"now":         "0 second ago",
		"ana":         "1 day ago",
		"ojo":         "day",
		"ope":         "december",
		"eti":         "friday",
		"gmt":         "gmt",
		"ola":         "in 1 day",
		"aje":         "monday",
		"osu":         "month",
		"utc":         "utc",
		"ose":         "week",
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
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ojo abameta|this minute|iseju aaya|last month|next month|ojo isegun|this month|last week|last year|next week|next year|osu agemo|osu ebibi|osu erele|osu erena|osu okudu|osu owara|osu owewe|this hour|this week|this year|ojo aiku|osu belu|osu igbe|osu ogun|osu sere|abameta|ojo aje|ojo eti|osu ope|isegun|wakati|agemo|ebibi|erele|erena|iseju|ojobo|ojoru|okudu|owara|owewe|aaro|aiku|belu|igbe|odun|ogun|osan|sere|aje|ana|eti|gmt|now|ojo|ola|oni|ope|ose|osu|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var yo_BJ_Locale = merge(&yo_Locale, LocaleData{
	Name:      "yo-BJ",
	DateOrder: "DMY",
	Translations: map[string]string{
		"ɔjɔ abamɛta": "saturday",
		"oshu erele":  "february",
		"oshu shɛrɛ":  "january",
		"oshu agɛmɔ":  "july",
		"oshu okudu":  "june",
		"oshu ɛrɛna":  "march",
		"oshu ɛbibi":  "may",
		"oshu ɔwara":  "october",
		"isɛju aaya":  "second",
		"oshu owewe":  "september",
		"ɔjɔ isɛgun":  "tuesday",
		"oshu igbe":   "april",
		"oshu ogun":   "august",
		"oshu belu":   "november",
		"oshu ɔpɛ":    "december",
		"ɔjɔ aiku":    "sunday",
		"ɔjɔ ɛti":     "friday",
		"ɔjɔ aje":     "monday",
		"abamɛta":     "saturday",
		"isɛgun":      "tuesday",
		"shɛrɛ":       "january",
		"agɛmɔ":       "july",
		"ɛrɛna":       "march",
		"ɛbibi":       "may",
		"isɛju":       "minute",
		"ɔwara":       "october",
		"ɔjɔbɔ":       "thursday",
		"ɔjɔru":       "wednesday",
		"aarɔ":        "am",
		"ɔsan":        "pm",
		"ɔdun":        "year",
		"ɔjɔ":         "day",
		"ɔpɛ":         "december",
		"ɛti":         "friday",
		"gmt":         "gmt",
		"ɔla":         "in 1 day",
		"utc":         "utc",
		"ɔse":         "week",
		"am":          "am",
		"pm":          "pm",
		" ":           " ",
		"+":           "+",
		"-":           "-",
		".":           ".",
		"/":           "/",
		":":           ":",
		"z":           "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ɔjɔ abamɛta|isɛju aaya|oshu agɛmɔ|oshu erele|oshu okudu|oshu owewe|oshu shɛrɛ|oshu ɔwara|oshu ɛbibi|oshu ɛrɛna|ɔjɔ isɛgun|oshu belu|oshu igbe|oshu ogun|oshu ɔpɛ|ɔjɔ aiku|abamɛta|ɔjɔ aje|ɔjɔ ɛti|isɛgun|agɛmɔ|isɛju|shɛrɛ|ɔjɔbɔ|ɔjɔru|ɔwara|ɛbibi|ɛrɛna|aarɔ|ɔdun|ɔsan|gmt|utc|ɔjɔ|ɔla|ɔpɛ|ɔse|ɛti|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

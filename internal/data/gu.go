// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var gu_Locale = merge(nil, LocaleData{
	Name:      "gu",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"આવતા અઠવાડિય": "in 1 week",
		"ગયા અઠવાડિય":  "1 week ago",
		"આ અઠવાડિય":    "0 week ago",
		"આવતા મહિન":    "in 1 month",
		"ગયા મહિન":     "1 month ago",
		"આવતા વરષ":     "in 1 year",
		"આ મિનિટ":      "0 minute ago",
		"ગયા વરષ":      "1 year ago",
		"આવતીકાલ":      "in 1 day",
		"જાનયઆરી":      "january",
		"અઠવાડિય":      "week",
		"આ કલાક":       "0 hour ago",
		"આ મહિન":       "0 month ago",
		"ડિસમબર":       "december",
		"ફબરઆરી":       "february",
		"શકરવાર":       "friday",
		"સોમવાર":       "monday",
		"ઑકટોબર":       "october",
		"શનિવાર":       "saturday",
		"સપટમબર":       "september",
		"રવિવાર":       "sunday",
		"મગળવાર":       "tuesday",
		"આ વરષ":        "0 year ago",
		"ગઈકાલ":        "1 day ago",
		"એપરિલ":        "april",
		"મિનિટ":        "minute",
		"મહિનો":        "month",
		"નવમબર":        "november",
		"ગરવાર":        "thursday",
		"બધવાર":        "wednesday",
		"હમણા":         "0 second ago",
		"ઑગસટ":         "august",
		"દિવસ":         "day",
		"કલાક":         "hour",
		"જાનય":         "january",
		"જલાઈ":         "july",
		"મારચ":         "march",
		"ઑકટો":         "october",
		"સકનડ":         "second",
		"ડિસ":          "december",
		"ફબર":          "february",
		"શકર":          "friday",
		"gmt":          "gmt",
		"સોમ":          "monday",
		"શનિ":          "saturday",
		"સપટ":          "september",
		"રવિ":          "sunday",
		"મગળ":          "tuesday",
		"utc":          "utc",
		"વરષ":          "year",
		"આજ":           "0 day ago",
		"am":           "am",
		"જન":           "june",
		"મિ":           "minute",
		"નવ":           "november",
		"pm":           "pm",
		"ગર":           "thursday",
		"બધ":           "wednesday",
		"અઠ":           "week",
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
		"ક":            "hour",
		"મ":            "may",
		"સ":            "second",
		"વ":            "year",
		"z":            "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) અઠવાડિયા પહલા`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) મિનિટ પહલા`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) મહિના પહલા`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) અઠવાડિયામા`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) દિવસ પહલા`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) કલાક પહલા`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) સકડ પહલા`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) વરષ પહલા`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) અઠ પહલા`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) મિનિટમા`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) મહિનામા`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) દિવસમા`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) કલાકમા`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) સકડમા`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) અઠ મા`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) વરષમા`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ અઠવાડિયા પહલા|\d+ અઠવાડિયામા|\d+ મહિના પહલા|\d+ મિનિટ પહલા|\d+ કલાક પહલા|\d+ દિવસ પહલા|\d+ વરષ પહલા|\d+ સકડ પહલા|\d+ અઠ પહલા|\d+ મહિનામા|\d+ મિનિટમા|\d+ કલાકમા|\d+ દિવસમા|\d+ અઠ મા|\d+ વરષમા|\d+ સકડમા)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ અઠવાડિયા પહલા|\d+ અઠવાડિયામા|\d+ મહિના પહલા|\d+ મિનિટ પહલા|\d+ કલાક પહલા|\d+ દિવસ પહલા|\d+ વરષ પહલા|\d+ સકડ પહલા|\d+ અઠ પહલા|\d+ મહિનામા|\d+ મિનિટમા|\d+ કલાકમા|\d+ દિવસમા|\d+ અઠ મા|\d+ વરષમા|\d+ સકડમા)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(આવતા અઠવાડિય|ગયા અઠવાડિય|આ અઠવાડિય|આવતા મહિન|આવતા વરષ|ગયા મહિન|અઠવાડિય|આ મિનિટ|આવતીકાલ|ગયા વરષ|જાનયઆરી|આ કલાક|આ મહિન|ઑકટોબર|ડિસમબર|ફબરઆરી|મગળવાર|રવિવાર|શકરવાર|શનિવાર|સપટમબર|સોમવાર|આ વરષ|એપરિલ|ગઈકાલ|ગરવાર|નવમબર|બધવાર|મહિનો|મિનિટ|ઑકટો|ઑગસટ|કલાક|જલાઈ|જાનય|દિવસ|મારચ|સકનડ|હમણા|gmt|utc|ડિસ|ફબર|મગળ|રવિ|વરષ|શકર|શનિ|સપટ|સોમ|\+|\.|\[|\]|\||am|pm|અઠ|આજ|ગર|જન|નવ|બધ|મિ| |'|,|-|/|:|;|@|z|ક|મ|વ|સ)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

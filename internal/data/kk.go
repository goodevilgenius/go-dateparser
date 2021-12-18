// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var kk_Locale = merge(nil, LocaleData{
	Name:      "kk",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"былтырғы жыл": "1 year ago",
		"келесі апта":  "in 1 week",
		"биылғы жыл":   "0 year ago",
		"өткен апта":   "1 week ago",
		"келесі жыл":   "in 1 year",
		"осы сағат":    "0 hour ago",
		"осы минут":    "0 minute ago",
		"желтоқсан":    "december",
		"келесі аи":    "in 1 month",
		"осы апта":     "0 week ago",
		"өткен аи":     "1 month ago",
		"дүисенбі":     "monday",
		"қыркүиек":     "september",
		"жексенбі":     "sunday",
		"беисенбі":     "thursday",
		"сеисенбі":     "tuesday",
		"сәрсенбі":     "wednesday",
		"осы аи":       "0 month ago",
		"қаңтар":       "january",
		"маусым":       "june",
		"наурыз":       "march",
		"қараша":       "november",
		"секунд":       "second",
		"бүгін":        "0 day ago",
		"қазір":        "0 second ago",
		"сәуір":        "april",
		"тамыз":        "august",
		"ақпан":        "february",
		"сағат":        "hour",
		"ертең":        "in 1 day",
		"шілде":        "july",
		"мамыр":        "may",
		"минут":        "minute",
		"қазан":        "october",
		"сенбі":        "saturday",
		"кеше":         "1 day ago",
		"жұма":         "friday",
		"апта":         "week",
		"сәу":          "april",
		"там":          "august",
		"күн":          "day",
		"жел":          "december",
		"ақп":          "february",
		"gmt":          "gmt",
		"сағ":          "hour",
		"қаң":          "january",
		"шіл":          "july",
		"мау":          "june",
		"нау":          "march",
		"мам":          "may",
		"мин":          "minute",
		"қар":          "november",
		"қаз":          "october",
		"қыр":          "september",
		"utc":          "utc",
		"жыл":          "year",
		"am":           "am",
		"жм":           "friday",
		"дс":           "monday",
		"аи":           "month",
		"pm":           "pm",
		"сб":           "saturday",
		"жс":           "sunday",
		"бс":           "thursday",
		"сс":           "tuesday",
		"ср":           "wednesday",
		"ап":           "week",
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
		"с":            "second",
		"ж":            "year",
		"z":            "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) секундтан кеиін`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) сағаттан кеиін`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) минуттан кеиін`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) аптадан кеиін`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) секунд бұрын`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) күннен кеиін`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) жылдан кеиін`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) сағат бұрын`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) минут бұрын`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) аидан кеиін`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) апта бұрын`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) күн бұрын`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) сағ бұрын`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) мин бұрын`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) сек бұрын`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) жыл бұрын`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) сағ кеиін`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) мин кеиін`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) сек кеиін`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) аи бұрын`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ап бұрын`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ап кеиін`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ж бұрын`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) ж кеиін`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ секундтан кеиін|\d+ минуттан кеиін|\d+ сағаттан кеиін|\d+ аптадан кеиін|\d+ жылдан кеиін|\d+ күннен кеиін|\d+ секунд бұрын|\d+ аидан кеиін|\d+ минут бұрын|\d+ сағат бұрын|\d+ апта бұрын|\d+ жыл бұрын|\d+ күн бұрын|\d+ мин бұрын|\d+ мин кеиін|\d+ сағ бұрын|\d+ сағ кеиін|\d+ сек бұрын|\d+ сек кеиін|\d+ аи бұрын|\d+ ап бұрын|\d+ ап кеиін|\d+ ж бұрын|\d+ ж кеиін)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ секундтан кеиін|\d+ минуттан кеиін|\d+ сағаттан кеиін|\d+ аптадан кеиін|\d+ жылдан кеиін|\d+ күннен кеиін|\d+ секунд бұрын|\d+ аидан кеиін|\d+ минут бұрын|\d+ сағат бұрын|\d+ апта бұрын|\d+ жыл бұрын|\d+ күн бұрын|\d+ мин бұрын|\d+ мин кеиін|\d+ сағ бұрын|\d+ сағ кеиін|\d+ сек бұрын|\d+ сек кеиін|\d+ аи бұрын|\d+ ап бұрын|\d+ ап кеиін|\d+ ж бұрын|\d+ ж кеиін)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(былтырғы жыл|келесі апта|биылғы жыл|келесі жыл|өткен апта|желтоқсан|келесі аи|осы минут|осы сағат|беисенбі|дүисенбі|жексенбі|осы апта|сеисенбі|сәрсенбі|қыркүиек|өткен аи|маусым|наурыз|осы аи|секунд|қараша|қаңтар|ақпан|бүгін|ертең|мамыр|минут|сағат|сенбі|сәуір|тамыз|шілде|қазан|қазір|апта|жұма|кеше|gmt|utc|ақп|жел|жыл|күн|мам|мау|мин|нау|сағ|сәу|там|шіл|қаз|қар|қаң|қыр|\+|\.|\[|\]|\||am|pm|аи|ап|бс|дс|жм|жс|сб|ср|сс| |'|,|-|/|:|;|@|z|ж|с)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

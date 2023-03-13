// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var dsb_Locale = merge(nil, LocaleData{
	Name:      "dsb",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuwyzóěłńśź`),
	Translations: map[string]string{
		"wotpołdnja": "pm",
		"dopołdnja":  "am",
		"ponjezele":  "monday",
		"september":  "september",
		"septembra":  "september",
		"december":   "december",
		"decembra":   "december",
		"februara":   "february",
		"nowember":   "november",
		"nowembra":   "november",
		"awgusta":    "august",
		"februar":    "february",
		"januara":    "january",
		"njezela":    "sunday",
		"oktober":    "october",
		"oktobra":    "october",
		"sekunda":    "second",
		"stwortk":    "thursday",
		"wałtora":    "tuesday",
		"apryla":     "april",
		"awgust":     "august",
		"gozina":     "hour",
		"januar":     "january",
		"julija":     "july",
		"junija":     "june",
		"minuta":     "minute",
		"mjasec":     "month",
		"sobota":     "saturday",
		"srjoda":     "wednesday",
		"apryl":      "april",
		"julij":      "july",
		"junij":      "june",
		"merca":      "march",
		"tyzen":      "week",
		"leto":       "year",
		"maja":       "may",
		"merc":       "march",
		"mjas":       "month",
		"petk":       "friday",
		"apr":        "april",
		"awg":        "august",
		"dec":        "december",
		"feb":        "february",
		"gmt":        "gmt",
		"goz":        "hour",
		"jan":        "january",
		"jul":        "july",
		"jun":        "june",
		"maj":        "may",
		"mer":        "march",
		"min":        "minute",
		"nje":        "sunday",
		"now":        "november",
		"okt":        "october",
		"pet":        "friday",
		"pon":        "monday",
		"sek":        "second",
		"sep":        "september",
		"sob":        "saturday",
		"srj":        "wednesday",
		"stw":        "thursday",
		"tyz":        "week",
		"utc":        "utc",
		"wał":        "tuesday",
		"zen":        "day",
		"am":         "am",
		"pm":         "pm",
		" ":          " ",
		"'":          "",
		"+":          "+",
		",":          "",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"g":          "hour",
		"l":          "year",
		"m":          "minute",
		"s":          "second",
		"z":          "z",
		"|":          "",
	},
	RelativeType: map[string]string{
		"psiducy mjasec": "in 1 month",
		"psiducy tyzen":  "in 1 week",
		"sledny mjasec":  "1 month ago",
		"sledny tyzen":   "1 week ago",
		"this minute":    "0 minute ago",
		"ten mjasec":     "0 month ago",
		"ten tyzen":      "0 week ago",
		"this hour":      "0 hour ago",
		"letosa":         "0 year ago",
		"witse":          "in 1 day",
		"zinsa":          "0 day ago",
		"znowa":          "in 1 year",
		"cora":           "1 day ago",
		"łoni":           "1 year ago",
		"now":            "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)psed (\d+) mjasecami`), "$1 month ago"},
		{regexp.MustCompile(`(?i)psed (\d+) sekundami`), "$1 second ago"},
		{regexp.MustCompile(`(?i)psed (\d+) tyzenjami`), "$1 week ago"},
		{regexp.MustCompile(`(?i)psed (\d+) gozinami`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)psed (\d+) minutami`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)psed (\d+) mjasecom`), "$1 month ago"},
		{regexp.MustCompile(`(?i)psed (\d+) tyzenjom`), "$1 week ago"},
		{regexp.MustCompile(`(?i)psed (\d+) sekundu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)psed (\d+) dnjami`), "$1 day ago"},
		{regexp.MustCompile(`(?i)psed (\d+) gozinu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)psed (\d+) letami`), "$1 year ago"},
		{regexp.MustCompile(`(?i)psed (\d+) minutu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)za (\d+) mjasecow`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) sekundow`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) tyzenjow`), "in $1 week"},
		{regexp.MustCompile(`(?i)psed (\d+) dnjom`), "$1 day ago"},
		{regexp.MustCompile(`(?i)psed (\d+) letom`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) minutow`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) sekundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)psed (\d+) mjas`), "$1 month ago"},
		{regexp.MustCompile(`(?i)za (\d+) gozinu`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) minutu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) mjasec`), "in $1 month"},
		{regexp.MustCompile(`(?i)psed (\d+) dnj`), "$1 day ago"},
		{regexp.MustCompile(`(?i)psed (\d+) goz`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)psed (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)psed (\d+) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)psed (\d+) tyz`), "$1 week ago"},
		{regexp.MustCompile(`(?i)za (\d+) dnjow`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) gozin`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) tyzen`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) leto`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) mjas`), "in $1 month"},
		{regexp.MustCompile(`(?i)psed (\d+) d`), "$1 day ago"},
		{regexp.MustCompile(`(?i)psed (\d+) g`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)psed (\d+) l`), "$1 year ago"},
		{regexp.MustCompile(`(?i)psed (\d+) m`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)psed (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) dnj`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) goz`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) let`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) tyz`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) zen`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) g`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) l`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) m`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) s`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) z`), "in $1 day"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(psed \d+ mjasecami|psed \d+ sekundami|psed \d+ tyzenjami|psed \d+ gozinami|psed \d+ minutami|psed \d+ mjasecom|psed \d+ tyzenjom|psed \d+ sekundu|psed \d+ dnjami|psed \d+ gozinu|psed \d+ letami|psed \d+ minutu|za \d+ mjasecow|za \d+ sekundow|za \d+ tyzenjow|psed \d+ dnjom|psed \d+ letom|za \d+ minutow|za \d+ sekundu|psed \d+ mjas|za \d+ gozinu|za \d+ minutu|za \d+ mjasec|psed \d+ dnj|psed \d+ goz|psed \d+ min|psed \d+ sek|psed \d+ tyz|za \d+ dnjow|za \d+ gozin|za \d+ tyzen|za \d+ leto|za \d+ mjas|psed \d+ d|psed \d+ g|psed \d+ l|psed \d+ m|psed \d+ s|za \d+ dnj|za \d+ goz|za \d+ let|za \d+ min|za \d+ sek|za \d+ tyz|za \d+ zen|za \d+ g|za \d+ l|za \d+ m|za \d+ s|za \d+ z)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(psed \d+ mjasecami|psed \d+ sekundami|psed \d+ tyzenjami|psed \d+ gozinami|psed \d+ minutami|psed \d+ mjasecom|psed \d+ tyzenjom|psed \d+ sekundu|psed \d+ dnjami|psed \d+ gozinu|psed \d+ letami|psed \d+ minutu|za \d+ mjasecow|za \d+ sekundow|za \d+ tyzenjow|psed \d+ dnjom|psed \d+ letom|za \d+ minutow|za \d+ sekundu|psed \d+ mjas|za \d+ gozinu|za \d+ minutu|za \d+ mjasec|psed \d+ dnj|psed \d+ goz|psed \d+ min|psed \d+ sek|psed \d+ tyz|za \d+ dnjow|za \d+ gozin|za \d+ tyzen|za \d+ leto|za \d+ mjas|psed \d+ d|psed \d+ g|psed \d+ l|psed \d+ m|psed \d+ s|za \d+ dnj|za \d+ goz|za \d+ let|za \d+ min|za \d+ sek|za \d+ tyz|za \d+ zen|za \d+ g|za \d+ l|za \d+ m|za \d+ s|za \d+ z)$`),
	KnownWords:      []string{"psiducy mjasec", "psiducy tyzen", "sledny mjasec", "sledny tyzen", "this minute", "ten mjasec", "wotpołdnja", "dopołdnja", "ponjezele", "september", "septembra", "ten tyzen", "this hour", "december", "decembra", "februara", "nowember", "nowembra", "awgusta", "februar", "januara", "njezela", "oktober", "oktobra", "sekunda", "stwortk", "wałtora", "apryla", "awgust", "gozina", "januar", "julija", "junija", "letosa", "minuta", "mjasec", "sobota", "srjoda", "apryl", "julij", "junij", "merca", "tyzen", "witse", "zinsa", "znowa", "cora", "leto", "maja", "merc", "mjas", "petk", "łoni", "apr", "awg", "dec", "feb", "gmt", "goz", "jan", "jul", "jun", "maj", "mer", "min", "nje", "now", "okt", "pet", "pon", "sek", "sep", "sob", "srj", "stw", "tyz", "utc", "wał", "zen", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "g", "l", "m", "s", "z", "|"},
})

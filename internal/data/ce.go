// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ce_Locale = merge(nil, LocaleData{
	Name:      "ce",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"рогӏерчу кӏирнахь": "in 1 week",
		"карарчу кӏирнахь":  "0 week ago",
		"даханчу кӏирнахь":  "1 week ago",
		"рогӏерчу баттахь":  "in 1 month",
		"карарчу баттахь":   "0 month ago",
		"баханчу баттахь":   "1 month ago",
		"рогӏерчу шарахь":   "in 1 year",
		"карарчу шарахь":    "0 year ago",
		"даханчу шарахь":    "1 year ago",
		"пӏераскан де":      "friday",
		"this minute":       "0 minute ago",
		"кӏиранан де":       "sunday",
		"оршотан де":        "monday",
		"шинарин де":        "tuesday",
		"кхаарин де":        "wednesday",
		"this hour":         "0 hour ago",
		"сентябрь":          "september",
		"еарин де":          "thursday",
		"селхана":           "1 day ago",
		"декабрь":           "december",
		"февраль":           "february",
		"октябрь":           "october",
		"тахана":            "0 day ago",
		"апрель":            "april",
		"август":            "august",
		"январь":            "january",
		"ноябрь":            "november",
		"шот де":            "saturday",
		"секунд":            "second",
		"сахьт":             "hour",
		"кхана":             "in 1 day",
		"минот":             "minute",
		"кӏира":             "week",
		"сахь":              "hour",
		"июль":              "july",
		"июнь":              "june",
		"март":              "march",
		"бутт":              "month",
		"кӏир":              "week",
		"now":               "0 second ago",
		"апр":               "april",
		"авг":               "august",
		"дек":               "december",
		"фев":               "february",
		"gmt":               "gmt",
		"янв":               "january",
		"июл":               "july",
		"июн":               "june",
		"мар":               "march",
		"маи":               "may",
		"мин":               "minute",
		"бут":               "month",
		"ноя":               "november",
		"окт":               "october",
		"сек":               "second",
		"сен":               "september",
		"utc":               "utc",
		"am":                "am",
		"де":                "day",
		"pm":                "pm",
		"шо":                "year",
		"'":                 "",
		",":                 "",
		";":                 "",
		"@":                 "",
		"[":                 "",
		"]":                 "",
		"|":                 "",
		" ":                 " ",
		"+":                 "+",
		"-":                 "-",
		".":                 ".",
		"/":                 "/",
		":":                 ":",
		"ш":                 "year",
		"z":                 "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) секунд хьалха`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) сахьт хьалха`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) минот хьалха`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) кӏира хьалха`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) сахьт даьлча`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) секунд яьлча`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) кӏира даьлча`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) сахь хьалха`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) бутт хьалха`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) кӏир хьалха`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) сахь даьлча`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) минот яьлча`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) бутт баьлча`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) кӏир даьлча`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) мин хьалха`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) сек хьалха`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) де хьалха`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) шо хьалха`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) де даьлча`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) мин яьлча`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) сек яьлча`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) шо даьлча`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) д хьалха`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) б хьалха`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ш хьалха`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) д даьлча`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) б баьлча`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) ш даьлча`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ секунд хьалха|\d+ кӏира даьлча|\d+ кӏира хьалха|\d+ минот хьалха|\d+ сахьт даьлча|\d+ сахьт хьалха|\d+ секунд яьлча|\d+ бутт баьлча|\d+ бутт хьалха|\d+ кӏир даьлча|\d+ кӏир хьалха|\d+ минот яьлча|\d+ сахь даьлча|\d+ сахь хьалха|\d+ мин хьалха|\d+ сек хьалха|\d+ де даьлча|\d+ де хьалха|\d+ мин яьлча|\d+ сек яьлча|\d+ шо даьлча|\d+ шо хьалха|\d+ б баьлча|\d+ б хьалха|\d+ д даьлча|\d+ д хьалха|\d+ ш даьлча|\d+ ш хьалха)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ секунд хьалха|\d+ кӏира даьлча|\d+ кӏира хьалха|\d+ минот хьалха|\d+ сахьт даьлча|\d+ сахьт хьалха|\d+ секунд яьлча|\d+ бутт баьлча|\d+ бутт хьалха|\d+ кӏир даьлча|\d+ кӏир хьалха|\d+ минот яьлча|\d+ сахь даьлча|\d+ сахь хьалха|\d+ мин хьалха|\d+ сек хьалха|\d+ де даьлча|\d+ де хьалха|\d+ мин яьлча|\d+ сек яьлча|\d+ шо даьлча|\d+ шо хьалха|\d+ б баьлча|\d+ б хьалха|\d+ д даьлча|\d+ д хьалха|\d+ ш даьлча|\d+ ш хьалха)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(рогӏерчу кӏирнахь|даханчу кӏирнахь|карарчу кӏирнахь|рогӏерчу баттахь|баханчу баттахь|карарчу баттахь|рогӏерчу шарахь|даханчу шарахь|карарчу шарахь|пӏераскан де|this minute|кӏиранан де|кхаарин де|оршотан де|шинарин де|this hour|еарин де|сентябрь|декабрь|октябрь|селхана|февраль|август|апрель|ноябрь|секунд|тахана|шот де|январь|кхана|кӏира|минот|сахьт|бутт|июль|июнь|кӏир|март|сахь|gmt|now|utc|авг|апр|бут|дек|июл|июн|маи|мар|мин|ноя|окт|сек|сен|фев|янв|\+|\.|\[|\]|\||am|pm|де|шо| |'|,|-|/|:|;|@|z|ш)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

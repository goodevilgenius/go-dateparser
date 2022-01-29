// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ro_Locale = merge(nil, LocaleData{
	Name:                  "ro",
	DateOrder:             "DMY",
	Charset:               []rune(`bcdefghijlnorstuvz`),
	SentenceSplitterGroup: 1,
	Translations: map[string]string{
		"septembrie": "september",
		"decembrie":  "december",
		"februarie":  "february",
		"noiembrie":  "november",
		"octombrie":  "october",
		"saptamana":  "week",
		"saptamani":  "week",
		"ianuarie":   "january",
		"duminica":   "sunday",
		"miercuri":   "wednesday",
		"in urma":    "ago",
		"aprilie":    "april",
		"sambata":    "saturday",
		"secunda":    "second",
		"secunde":    "second",
		"august":     "august",
		"vineri":     "friday",
		"martie":     "march",
		"minute":     "minute",
		"iulie":      "july",
		"iunie":      "june",
		"minut":      "minute",
		"noiem":      "november",
		"marti":      "tuesday",
		"zile":       "day",
		"febr":       "february",
		"mart":       "march",
		"luna":       "month",
		"luni":       "month",
		"sept":       "september",
		"sapt":       "week",
		"apr":        "april",
		"aug":        "august",
		"dec":        "december",
		"feb":        "february",
		"vin":        "friday",
		"gmt":        "gmt",
		"ora":        "hour",
		"ore":        "hour",
		"ian":        "january",
		"iul":        "july",
		"iun":        "june",
		"mar":        "march",
		"mai":        "may",
		"min":        "minute",
		"lun":        "monday",
		"nov":        "november",
		"oct":        "october",
		"sam":        "saturday",
		"sec":        "second",
		"sep":        "september",
		"dum":        "sunday",
		"joi":        "thursday",
		"utc":        "utc",
		"mie":        "wednesday",
		"ani":        "year",
		"de":         "",
		"la":         "",
		"am":         "am",
		"zi":         "day",
		"in":         "in",
		"pm":         "pm",
		"mi":         "wednesday",
		"an":         "year",
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
		"h":          "hour",
		"m":          "minute",
		"s":          "second",
		"z":          "z",
	},
	RelativeType: map[string]string{
		"saptamana viitoare": "in 1 week",
		"saptamana aceasta":  "0 week ago",
		"saptamana trecuta":  "1 week ago",
		"minutul acesta":     "0 minute ago",
		"luna viitoare":      "in 1 month",
		"luna aceasta":       "0 month ago",
		"luna trecuta":       "1 month ago",
		"ora aceasta":        "0 hour ago",
		"anul acesta":        "0 year ago",
		"anul trecut":        "1 year ago",
		"anul viitor":        "in 1 year",
		"maine":              "in 1 day",
		"acum":               "0 second ago",
		"ieri":               "1 day ago",
		"azi":                "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)peste (\d+) de saptamani`), "in $1 week"},
		{regexp.MustCompile(`(?i)acum (\d+) de saptamani`), "$1 week ago"},
		{regexp.MustCompile(`(?i)peste (\d+) de secunde`), "in $1 second"},
		{regexp.MustCompile(`(?i)acum (\d+) de secunde`), "$1 second ago"},
		{regexp.MustCompile(`(?i)peste (\d+) de minute`), "in $1 minute"},
		{regexp.MustCompile(`(?i)peste (\d+) saptamana`), "in $1 week"},
		{regexp.MustCompile(`(?i)acum (\d+) de minute`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)acum (\d+) saptamana`), "$1 week ago"},
		{regexp.MustCompile(`(?i)peste (\d+) de zile`), "in $1 day"},
		{regexp.MustCompile(`(?i)peste (\d+) de luni`), "in $1 month"},
		{regexp.MustCompile(`(?i)peste (\d+) secunda`), "in $1 second"},
		{regexp.MustCompile(`(?i)acum (\d+) de zile`), "$1 day ago"},
		{regexp.MustCompile(`(?i)acum (\d+) de luni`), "$1 month ago"},
		{regexp.MustCompile(`(?i)acum (\d+) secunda`), "$1 second ago"},
		{regexp.MustCompile(`(?i)peste (\d+) de ore`), "in $1 hour"},
		{regexp.MustCompile(`(?i)peste (\d+) de ani`), "in $1 year"},
		{regexp.MustCompile(`(?i)acum (\d+) de ore`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)acum (\d+) de ani`), "$1 year ago"},
		{regexp.MustCompile(`(?i)peste (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)acum (\d+) minut`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)peste (\d+) zile`), "in $1 day"},
		{regexp.MustCompile(`(?i)peste (\d+) luna`), "in $1 month"},
		{regexp.MustCompile(`(?i)peste (\d+) luni`), "in $1 month"},
		{regexp.MustCompile(`(?i)peste (\d+) sapt`), "in $1 week"},
		{regexp.MustCompile(`(?i)acum (\d+) zile`), "$1 day ago"},
		{regexp.MustCompile(`(?i)acum (\d+) luna`), "$1 month ago"},
		{regexp.MustCompile(`(?i)acum (\d+) luni`), "$1 month ago"},
		{regexp.MustCompile(`(?i)acum (\d+) sapt`), "$1 week ago"},
		{regexp.MustCompile(`(?i)peste (\d+) ora`), "in $1 hour"},
		{regexp.MustCompile(`(?i)peste (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)peste (\d+) sec`), "in $1 second"},
		{regexp.MustCompile(`(?i)peste (\d+) ani`), "in $1 year"},
		{regexp.MustCompile(`(?i)acum (\d+) ora`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)acum (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)acum (\d+) sec`), "$1 second ago"},
		{regexp.MustCompile(`(?i)peste (\d+) zi`), "in $1 day"},
		{regexp.MustCompile(`(?i)peste (\d+) an`), "in $1 year"},
		{regexp.MustCompile(`(?i)acum (\d+) zi`), "$1 day ago"},
		{regexp.MustCompile(`(?i)acum (\d+) an`), "$1 year ago"},
		{regexp.MustCompile(`(?i)peste (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)acum (\d+) h`), "$1 hour ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(peste \d+ de saptamani|acum \d+ de saptamani|peste \d+ de secunde|acum \d+ de secunde|peste \d+ de minute|peste \d+ saptamana|acum \d+ de minute|acum \d+ saptamana|peste \d+ de luni|peste \d+ de zile|peste \d+ secunda|acum \d+ de luni|acum \d+ de zile|acum \d+ secunda|peste \d+ de ani|peste \d+ de ore|acum \d+ de ani|acum \d+ de ore|peste \d+ minut|acum \d+ minut|peste \d+ luna|peste \d+ luni|peste \d+ sapt|peste \d+ zile|acum \d+ luna|acum \d+ luni|acum \d+ sapt|acum \d+ zile|peste \d+ ani|peste \d+ min|peste \d+ ora|peste \d+ sec|acum \d+ min|acum \d+ ora|acum \d+ sec|peste \d+ an|peste \d+ zi|acum \d+ an|acum \d+ zi|peste \d+ h|acum \d+ h)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(peste \d+ de saptamani|acum \d+ de saptamani|peste \d+ de secunde|acum \d+ de secunde|peste \d+ de minute|peste \d+ saptamana|acum \d+ de minute|acum \d+ saptamana|peste \d+ de luni|peste \d+ de zile|peste \d+ secunda|acum \d+ de luni|acum \d+ de zile|acum \d+ secunda|peste \d+ de ani|peste \d+ de ore|acum \d+ de ani|acum \d+ de ore|peste \d+ minut|acum \d+ minut|peste \d+ luna|peste \d+ luni|peste \d+ sapt|peste \d+ zile|acum \d+ luna|acum \d+ luni|acum \d+ sapt|acum \d+ zile|peste \d+ ani|peste \d+ min|peste \d+ ora|peste \d+ sec|acum \d+ min|acum \d+ ora|acum \d+ sec|peste \d+ an|peste \d+ zi|acum \d+ an|acum \d+ zi|peste \d+ h|acum \d+ h)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(saptamana viitoare|saptamana aceasta|saptamana trecuta|minutul acesta|luna viitoare|luna aceasta|luna trecuta|anul acesta|anul trecut|anul viitor|ora aceasta|septembrie|decembrie|februarie|noiembrie|octombrie|saptamana|saptamani|duminica|ianuarie|miercuri|aprilie|in urma|sambata|secunda|secunde|august|martie|minute|vineri|iulie|iunie|maine|marti|minut|noiem|acum|febr|ieri|luna|luni|mart|sapt|sept|zile|ani|apr|aug|azi|dec|dum|feb|gmt|ian|iul|iun|joi|lun|mai|mar|mie|min|nov|oct|ora|ore|sam|sec|sep|utc|vin|\+|\.|\[|\]|\||am|an|de|in|la|mi|pm|zi| |'|,|-|/|:|;|@|h|m|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ro_MD_Locale = merge(&ro_Locale, LocaleData{
	Name:            "ro-MD",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(peste \d+ de saptamani|acum \d+ de saptamani|peste \d+ de secunde|acum \d+ de secunde|peste \d+ de minute|peste \d+ saptamana|acum \d+ de minute|acum \d+ saptamana|peste \d+ de luni|peste \d+ de zile|peste \d+ secunda|acum \d+ de luni|acum \d+ de zile|acum \d+ secunda|peste \d+ de ani|peste \d+ de ore|acum \d+ de ani|acum \d+ de ore|peste \d+ minut|acum \d+ minut|peste \d+ luna|peste \d+ luni|peste \d+ sapt|peste \d+ zile|acum \d+ luna|acum \d+ luni|acum \d+ sapt|acum \d+ zile|peste \d+ ani|peste \d+ min|peste \d+ ora|peste \d+ sec|acum \d+ min|acum \d+ ora|acum \d+ sec|peste \d+ an|peste \d+ zi|acum \d+ an|acum \d+ zi|peste \d+ h|acum \d+ h)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(peste \d+ de saptamani|acum \d+ de saptamani|peste \d+ de secunde|acum \d+ de secunde|peste \d+ de minute|peste \d+ saptamana|acum \d+ de minute|acum \d+ saptamana|peste \d+ de luni|peste \d+ de zile|peste \d+ secunda|acum \d+ de luni|acum \d+ de zile|acum \d+ secunda|peste \d+ de ani|peste \d+ de ore|acum \d+ de ani|acum \d+ de ore|peste \d+ minut|acum \d+ minut|peste \d+ luna|peste \d+ luni|peste \d+ sapt|peste \d+ zile|acum \d+ luna|acum \d+ luni|acum \d+ sapt|acum \d+ zile|peste \d+ ani|peste \d+ min|peste \d+ ora|peste \d+ sec|acum \d+ min|acum \d+ ora|acum \d+ sec|peste \d+ an|peste \d+ zi|acum \d+ an|acum \d+ zi|peste \d+ h|acum \d+ h)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(saptamana viitoare|saptamana aceasta|saptamana trecuta|minutul acesta|luna viitoare|luna aceasta|luna trecuta|anul acesta|anul trecut|anul viitor|ora aceasta|septembrie|decembrie|februarie|noiembrie|octombrie|saptamana|saptamani|duminica|ianuarie|miercuri|aprilie|in urma|sambata|secunda|secunde|august|martie|minute|vineri|iulie|iunie|maine|marti|minut|noiem|acum|febr|ieri|luna|luni|mart|sapt|sept|zile|ani|apr|aug|azi|dec|dum|feb|gmt|ian|iul|iun|joi|lun|mai|mar|mie|min|nov|oct|ora|ore|sam|sec|sep|utc|vin|\+|\.|\[|\]|\||am|an|de|in|la|mi|pm|zi| |'|,|-|/|:|;|@|h|m|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

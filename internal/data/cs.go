// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var cs_Locale = merge(nil, LocaleData{
	Name:                  "cs",
	DateOrder:             "DMY",
	Charset:               []rune(`bcdeghijklnorstuvyzáíúýčěřšů`),
	SentenceSplitterGroup: 1,
	Translations: map[string]string{
		"listopadu": "november",
		"priblizne": "",
		"cervence":  "july",
		"cervenec":  "july",
		"hodinami":  "hour",
		"listopad":  "november",
		"prosince":  "december",
		"prosinec":  "december",
		"ctvrtek":   "thursday",
		"pondeli":   "monday",
		"sekunda":   "second",
		"sekundy":   "second",
		"vterina":   "second",
		"vteriny":   "second",
		"brezen":    "march",
		"brezna":    "march",
		"cerven":    "june",
		"cervna":    "june",
		"hodina":    "hour",
		"hodinu":    "hour",
		"hodiny":    "hour",
		"kveten":    "may",
		"kvetna":    "may",
		"mesice":    "month",
		"mesicu":    "month",
		"minuta":    "minute",
		"minuty":    "minute",
		"nedele":    "sunday",
		"nedeli":    "sunday",
		"sekund":    "second",
		"sobota":    "saturday",
		"sobotu":    "saturday",
		"streda":    "wednesday",
		"stredu":    "wednesday",
		"vterin":    "second",
		"duben":     "april",
		"dubna":     "april",
		"hodin":     "hour",
		"leden":     "january",
		"ledna":     "january",
		"mesic":     "month",
		"minut":     "minute",
		"patek":     "friday",
		"rijen":     "october",
		"rijna":     "october",
		"srpen":     "august",
		"srpna":     "august",
		"tyden":     "week",
		"tydnu":     "week",
		"tydny":     "week",
		"unora":     "february",
		"utery":     "tuesday",
		"pred":      "ago",
		"roku":      "year",
		"roky":      "year",
		"unor":      "february",
		"zari":      "september",
		"bre":       "march",
		"cer":       "june",
		"crc":       "july",
		"ctv":       "thursday",
		"cvc":       "july",
		"cvn":       "june",
		"den":       "day",
		"dnu":       "day",
		"dny":       "day",
		"dop":       "am",
		"dub":       "april",
		"gmt":       "gmt",
		"kve":       "may",
		"led":       "january",
		"lis":       "november",
		"mes":       "month",
		"min":       "minute",
		"ned":       "sunday",
		"odp":       "pm",
		"pat":       "friday",
		"pon":       "monday",
		"pro":       "december",
		"rij":       "october",
		"rok":       "year",
		"sob":       "saturday",
		"srp":       "august",
		"str":       "wednesday",
		"tyd":       "week",
		"uno":       "february",
		"utc":       "utc",
		"ute":       "tuesday",
		"zar":       "september",
		"am":        "am",
		"ct":        "thursday",
		"ne":        "sunday",
		"pa":        "friday",
		"pm":        "pm",
		"po":        "monday",
		"so":        "saturday",
		"st":        "wednesday",
		"ut":        "tuesday",
		"ve":        "in",
		" ":         " ",
		"'":         "",
		"+":         "+",
		",":         "",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"h":         "hour",
		"r":         "year",
		"s":         "second",
		"v":         "in",
		"z":         "z",
		"|":         "",
	},
	RelativeType: map[string]string{
		"minuly mesic": "1 month ago",
		"minuly tyden": "1 week ago",
		"pristi mesic": "in 1 month",
		"pristi tyden": "in 1 week",
		"predevcirem":  "2 day ago",
		"tento mesic":  "0 month ago",
		"tento tyden":  "0 week ago",
		"tuto hodinu":  "0 hour ago",
		"tuto minutu":  "0 minute ago",
		"minuly rok":   "1 year ago",
		"minuly tyd":   "1 week ago",
		"pristi rok":   "in 1 year",
		"pristi tyd":   "in 1 week",
		"tento rok":    "0 year ago",
		"tento tyd":    "0 week ago",
		"vcera":        "1 day ago",
		"zitra":        "in 1 day",
		"dnes":         "0 day ago",
		"nyni":         "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)pred (\d+) sekundami`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) hodinami`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) minutami`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) sekundou`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) hodinou`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mesicem`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) minutou`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mesici`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tydnem`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pred (\d+) rokem`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tydny`), "$1 week ago"},
		{regexp.MustCompile(`(?i)za (\d+) sekundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)pred (\d+) dnem`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) lety`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) hodinu`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) mesicu`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) minutu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) sekund`), "in $1 second"},
		{regexp.MustCompile(`(?i)pred (\d+) dny`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mes`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tyd`), "$1 week ago"},
		{regexp.MustCompile(`(?i)za (\d+) hodin`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) mesic`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) tyden`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) tydnu`), "in $1 week"},
		{regexp.MustCompile(`(?i)pred (\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) l`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pred (\d+) r`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pred (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) den`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) dni`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) let`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) rok`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) tyd`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) l`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) r`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pred \d+ sekundami|pred \d+ hodinami|pred \d+ minutami|pred \d+ sekundou|pred \d+ hodinou|pred \d+ mesicem|pred \d+ minutou|pred \d+ mesici|pred \d+ tydnem|pred \d+ rokem|pred \d+ tydny|za \d+ sekundu|pred \d+ dnem|pred \d+ lety|za \d+ hodinu|za \d+ mesicu|za \d+ minutu|za \d+ sekund|pred \d+ dny|pred \d+ mes|pred \d+ min|pred \d+ tyd|za \d+ hodin|za \d+ mesic|za \d+ minut|za \d+ tyden|za \d+ tydnu|pred \d+ h|pred \d+ l|pred \d+ r|pred \d+ s|za \d+ den|za \d+ dni|za \d+ let|za \d+ mes|za \d+ min|za \d+ rok|za \d+ tyd|za \d+ h|za \d+ l|za \d+ r|za \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(pred \d+ sekundami|pred \d+ hodinami|pred \d+ minutami|pred \d+ sekundou|pred \d+ hodinou|pred \d+ mesicem|pred \d+ minutou|pred \d+ mesici|pred \d+ tydnem|pred \d+ rokem|pred \d+ tydny|za \d+ sekundu|pred \d+ dnem|pred \d+ lety|za \d+ hodinu|za \d+ mesicu|za \d+ minutu|za \d+ sekund|pred \d+ dny|pred \d+ mes|pred \d+ min|pred \d+ tyd|za \d+ hodin|za \d+ mesic|za \d+ minut|za \d+ tyden|za \d+ tydnu|pred \d+ h|pred \d+ l|pred \d+ r|pred \d+ s|za \d+ den|za \d+ dni|za \d+ let|za \d+ mes|za \d+ min|za \d+ rok|za \d+ tyd|za \d+ h|za \d+ l|za \d+ r|za \d+ s)$`),
	KnownWords:      []string{"minuly mesic", "minuly tyden", "pristi mesic", "pristi tyden", "predevcirem", "tento mesic", "tento tyden", "tuto hodinu", "tuto minutu", "minuly rok", "minuly tyd", "pristi rok", "pristi tyd", "listopadu", "priblizne", "tento rok", "tento tyd", "cervence", "cervenec", "hodinami", "listopad", "prosince", "prosinec", "ctvrtek", "pondeli", "sekunda", "sekundy", "vterina", "vteriny", "brezen", "brezna", "cerven", "cervna", "hodina", "hodinu", "hodiny", "kveten", "kvetna", "mesice", "mesicu", "minuta", "minuty", "nedele", "nedeli", "sekund", "sobota", "sobotu", "streda", "stredu", "vterin", "duben", "dubna", "hodin", "leden", "ledna", "mesic", "minut", "patek", "rijen", "rijna", "srpen", "srpna", "tyden", "tydnu", "tydny", "unora", "utery", "vcera", "zitra", "dnes", "nyni", "pred", "roku", "roky", "unor", "zari", "bre", "cer", "crc", "ctv", "cvc", "cvn", "den", "dnu", "dny", "dop", "dub", "gmt", "kve", "led", "lis", "mes", "min", "ned", "odp", "pat", "pon", "pro", "rij", "rok", "sob", "srp", "str", "tyd", "uno", "utc", "ute", "zar", "am", "ct", "ne", "pa", "pm", "po", "so", "st", "ut", "ve", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "h", "r", "s", "v", "z", "|"},
})

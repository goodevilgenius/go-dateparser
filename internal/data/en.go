// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var en_Locale = merge(nil, LocaleData{
	Name:         "en",
	DateOrder:    "MDY",
	SkipWords:    []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "about", "ad", "and", "at", "by", "just", "m", "nd", "of", "on", "rd", "st", "th", "the", "|"},
	PertainWords: []string{"of"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)less than 1 minute ago(\z|[^\pL\pM\d]|_)`), "${1}45 second ago${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(?:12\s+)?midnight(\z|[^\pL\pM\d]|_)`), "${1}00:00${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(?:12\s+)?noon(\z|[^\pL\pM\d]|_)`), "${1}12:00${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+)h(\d+)m?(\z|[^\pL\pM\d]|_)`), "${1}${2}:${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(from\s+)now(\z|[^\pL\pM\d]|_)`), "${1}${2}in${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)eleven(\z|[^\pL\pM\d]|_)`), "${1}11${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)twelve(\z|[^\pL\pM\d]|_)`), "${1}12${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)three(\z|[^\pL\pM\d]|_)`), "${1}3${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)seven(\z|[^\pL\pM\d]|_)`), "${1}7${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)eight(\z|[^\pL\pM\d]|_)`), "${1}8${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)four(\z|[^\pL\pM\d]|_)`), "${1}4${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)five(\z|[^\pL\pM\d]|_)`), "${1}5${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)nine(\z|[^\pL\pM\d]|_)`), "${1}9${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)one(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ten(\z|[^\pL\pM\d]|_)`), "${1}10${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)two(\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)six(\z|[^\pL\pM\d]|_)`), "${1}6${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)an(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)a(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
	},
	Translations: map[string]string{
		"day before yesterday": "2 day ago",
		"this minute":          "0 minute ago",
		"last decade":          "1 decade ago",
		"this decade":          "1 decade ago",
		"next decade":          "in 1 decade",
		"this month":           "0 month ago",
		"last month":           "1 month ago",
		"next month":           "in 1 month",
		"this hour":            "0 hour ago",
		"this week":            "0 week ago",
		"this year":            "0 year ago",
		"yesterday":            "1 day ago",
		"last week":            "1 week ago",
		"last year":            "1 year ago",
		"next week":            "in 1 week",
		"next year":            "in 1 year",
		"september":            "september",
		"wednesday":            "wednesday",
		"december":             "december",
		"february":             "february",
		"from now":             "in",
		"tomorrow":             "in 1 day",
		"november":             "november",
		"saturday":             "saturday",
		"thursday":             "thursday",
		"this mo":              "0 month ago",
		"this wk":              "0 week ago",
		"this yr":              "0 year ago",
		"last mo":              "1 month ago",
		"last wk":              "1 week ago",
		"last yr":              "1 year ago",
		"decades":              "decade",
		"next mo":              "in 1 month",
		"next wk":              "in 1 week",
		"next yr":              "in 1 year",
		"january":              "january",
		"minutes":              "minute",
		"october":              "october",
		"seconds":              "second",
		"tuesday":              "tuesday",
		"august":               "august",
		"decade":               "decade",
		"friday":               "friday",
		"minute":               "minute",
		"monday":               "monday",
		"months":               "month",
		"second":               "second",
		"sunday":               "sunday",
		"about":                "",
		"today":                "0 day ago",
		"april":                "april",
		"hours":                "hour",
		"march":                "march",
		"month":                "month",
		"weeks":                "week",
		"years":                "year",
		"just":                 "",
		"days":                 "day",
		"hour":                 "hour",
		"july":                 "july",
		"june":                 "june",
		"mins":                 "minute",
		"secs":                 "second",
		"sept":                 "september",
		"tues":                 "tuesday",
		"week":                 "week",
		"year":                 "year",
		"and":                  "",
		"the":                  "",
		"now":                  "0 second ago",
		"ago":                  "ago",
		"apr":                  "april",
		"aug":                  "august",
		"day":                  "day",
		"dec":                  "december",
		"feb":                  "february",
		"fri":                  "friday",
		"gmt":                  "gmt",
		"hrs":                  "hour",
		"jan":                  "january",
		"jul":                  "july",
		"jun":                  "june",
		"mar":                  "march",
		"may":                  "may",
		"min":                  "minute",
		"mon":                  "monday",
		"nov":                  "november",
		"oct":                  "october",
		"sat":                  "saturday",
		"sec":                  "second",
		"sep":                  "september",
		"sun":                  "sunday",
		"thu":                  "thursday",
		"tue":                  "tuesday",
		"utc":                  "utc",
		"wed":                  "wednesday",
		"ad":                   "",
		"at":                   "",
		"by":                   "",
		"nd":                   "",
		"of":                   "",
		"on":                   "",
		"rd":                   "",
		"st":                   "",
		"th":                   "",
		"am":                   "am",
		"hr":                   "hour",
		"in":                   "in",
		"mo":                   "month",
		"pm":                   "pm",
		"wk":                   "week",
		"yr":                   "year",
		"'":                    "",
		",":                    "",
		";":                    "",
		"@":                    "",
		"[":                    "",
		"]":                    "",
		"m":                    "",
		"|":                    "",
		" ":                    " ",
		"+":                    "+",
		"-":                    "-",
		".":                    ".",
		"/":                    "/",
		":":                    ":",
		"h":                    "hour",
		"s":                    "second",
		"z":                    "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) decades? ago`), "$1 decade ago"},
		{regexp.MustCompile(`(?i)(\d+) minutes ago`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) seconds ago`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+) decades?`), "in $1 decade"},
		{regexp.MustCompile(`(?i)(\d+) minute ago`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) months ago`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) second ago`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+) minutes`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+) seconds`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) hours ago`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) month ago`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) weeks ago`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) years ago`), "$1 year ago"},
		{regexp.MustCompile(`(?i)in (\d+) minute`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+) months`), "in $1 month"},
		{regexp.MustCompile(`(?i)in (\d+) second`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) days ago`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) hour ago`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) week ago`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) year ago`), "$1 year ago"},
		{regexp.MustCompile(`(?i)in (\d+) hours`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+) month`), "in $1 month"},
		{regexp.MustCompile(`(?i)in (\d+) weeks`), "in $1 week"},
		{regexp.MustCompile(`(?i)in (\d+) years`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) day ago`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) min ago`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) sec ago`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+) days`), "in $1 day"},
		{regexp.MustCompile(`(?i)in (\d+) hour`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+) week`), "in $1 week"},
		{regexp.MustCompile(`(?i)in (\d+) year`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) hr ago`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) mo ago`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) wk ago`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) yr ago`), "$1 year ago"},
		{regexp.MustCompile(`(?i)in (\d+) day`), "in $1 day"},
		{regexp.MustCompile(`(?i)in (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+) sec`), "in $1 second"},
		{regexp.MustCompile(`(?i)in (\d+) hr`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+) mo`), "in $1 month"},
		{regexp.MustCompile(`(?i)in (\d+) wk`), "in $1 week"},
		{regexp.MustCompile(`(?i)in (\d+) yr`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ decades? ago|\d+ minutes ago|\d+ seconds ago|in \d+ decades?|\d+ minute ago|\d+ months ago|\d+ second ago|in \d+ minutes|in \d+ seconds|\d+ hours ago|\d+ month ago|\d+ weeks ago|\d+ years ago|in \d+ minute|in \d+ months|in \d+ second|\d+ days ago|\d+ hour ago|\d+ week ago|\d+ year ago|in \d+ hours|in \d+ month|in \d+ weeks|in \d+ years|\d+ day ago|\d+ min ago|\d+ sec ago|in \d+ days|in \d+ hour|in \d+ week|in \d+ year|\d+ hr ago|\d+ mo ago|\d+ wk ago|\d+ yr ago|in \d+ day|in \d+ min|in \d+ sec|in \d+ hr|in \d+ mo|in \d+ wk|in \d+ yr)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ decades? ago|\d+ minutes ago|\d+ seconds ago|in \d+ decades?|\d+ minute ago|\d+ months ago|\d+ second ago|in \d+ minutes|in \d+ seconds|\d+ hours ago|\d+ month ago|\d+ weeks ago|\d+ years ago|in \d+ minute|in \d+ months|in \d+ second|\d+ days ago|\d+ hour ago|\d+ week ago|\d+ year ago|in \d+ hours|in \d+ month|in \d+ weeks|in \d+ years|\d+ day ago|\d+ min ago|\d+ sec ago|in \d+ days|in \d+ hour|in \d+ week|in \d+ year|\d+ hr ago|\d+ mo ago|\d+ wk ago|\d+ yr ago|in \d+ day|in \d+ min|in \d+ sec|in \d+ hr|in \d+ mo|in \d+ wk|in \d+ yr)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(day before yesterday|last decade|next decade|this decade|this minute|last month|next month|this month|last week|last year|next week|next year|september|this hour|this week|this year|wednesday|yesterday|december|february|from now|november|saturday|thursday|tomorrow|decades|january|last mo|last wk|last yr|minutes|next mo|next wk|next yr|october|seconds|this mo|this wk|this yr|tuesday|august|decade|friday|minute|monday|months|second|sunday|about|april|hours|march|month|today|weeks|years|days|hour|july|june|just|mins|secs|sept|tues|week|year|ago|and|apr|aug|day|dec|feb|fri|gmt|hrs|jan|jul|jun|mar|may|min|mon|nov|now|oct|sat|sec|sep|sun|the|thu|tue|utc|wed|\+|\.|\[|\]|\||ad|am|at|by|hr|in|mo|nd|of|on|pm|rd|st|th|wk|yr| |'|,|-|/|:|;|@|h|m|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var en_001_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-001",
	DateOrder: "DMY",
})

var en_150_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-150",
	DateOrder: "DMY",
})

var en_AG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-AG",
	DateOrder: "DMY",
})

var en_AI_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-AI",
	DateOrder: "DMY",
})

var en_AS_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-AS",
	DateOrder: "MDY",
})

var en_AT_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-AT",
	DateOrder: "DMY",
})

var en_AU_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-AU",
	DateOrder: "DMY",
})

var en_BB_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-BB",
	DateOrder: "DMY",
})

var en_BE_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-BE",
	DateOrder: "DMY",
})

var en_BI_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-BI",
	DateOrder: "MDY",
})

var en_BM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-BM",
	DateOrder: "DMY",
})

var en_BS_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-BS",
	DateOrder: "DMY",
})

var en_BW_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-BW",
	DateOrder: "DMY",
})

var en_BZ_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-BZ",
	DateOrder: "DMY",
})

var en_CA_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-CA",
	DateOrder: "YMD",
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) mins ago`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) secs ago`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) hrs ago`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) mos ago`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) wks ago`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) yrs ago`), "$1 year ago"},
		{regexp.MustCompile(`(?i)in (\d+) mins`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+) secs`), "in $1 second"},
		{regexp.MustCompile(`(?i)in (\d+) hrs`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+) mos`), "in $1 month"},
		{regexp.MustCompile(`(?i)in (\d+) wks`), "in $1 week"},
		{regexp.MustCompile(`(?i)in (\d+) yrs`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ mins ago|\d+ secs ago|\d+ hrs ago|\d+ mos ago|\d+ wks ago|\d+ yrs ago|in \d+ mins|in \d+ secs|in \d+ hrs|in \d+ mos|in \d+ wks|in \d+ yrs)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ mins ago|\d+ secs ago|\d+ hrs ago|\d+ mos ago|\d+ wks ago|\d+ yrs ago|in \d+ mins|in \d+ secs|in \d+ hrs|in \d+ mos|in \d+ wks|in \d+ yrs)$`),
})

var en_CC_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-CC",
	DateOrder: "DMY",
})

var en_CH_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-CH",
	DateOrder: "DMY",
})

var en_CK_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-CK",
	DateOrder: "DMY",
})

var en_CM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-CM",
	DateOrder: "DMY",
})

var en_CX_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-CX",
	DateOrder: "DMY",
})

var en_CY_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-CY",
	DateOrder: "DMY",
})

var en_DE_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-DE",
	DateOrder: "DMY",
})

var en_DG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-DG",
	DateOrder: "DMY",
})

var en_DK_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-DK",
	DateOrder: "DMY",
})

var en_DM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-DM",
	DateOrder: "DMY",
})

var en_ER_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-ER",
	DateOrder: "DMY",
})

var en_FI_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-FI",
	DateOrder: "DMY",
})

var en_FJ_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-FJ",
	DateOrder: "DMY",
})

var en_FK_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-FK",
	DateOrder: "DMY",
})

var en_FM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-FM",
	DateOrder: "DMY",
})

var en_GB_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-GB",
	DateOrder: "DMY",
})

var en_GD_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-GD",
	DateOrder: "DMY",
})

var en_GG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-GG",
	DateOrder: "DMY",
})

var en_GH_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-GH",
	DateOrder: "DMY",
})

var en_GI_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-GI",
	DateOrder: "DMY",
})

var en_GM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-GM",
	DateOrder: "DMY",
})

var en_GU_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-GU",
	DateOrder: "MDY",
})

var en_GY_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-GY",
	DateOrder: "DMY",
})

var en_HK_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-HK",
	DateOrder: "DMY",
})

var en_IE_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-IE",
	DateOrder: "DMY",
})

var en_IL_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-IL",
	DateOrder: "DMY",
})

var en_IM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-IM",
	DateOrder: "DMY",
})

var en_IN_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-IN",
	DateOrder: "DMY",
})

var en_IO_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-IO",
	DateOrder: "DMY",
})

var en_JE_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-JE",
	DateOrder: "DMY",
})

var en_JM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-JM",
	DateOrder: "DMY",
})

var en_KE_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-KE",
	DateOrder: "DMY",
})

var en_KI_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-KI",
	DateOrder: "DMY",
})

var en_KN_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-KN",
	DateOrder: "DMY",
})

var en_KY_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-KY",
	DateOrder: "DMY",
})

var en_LC_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-LC",
	DateOrder: "DMY",
})

var en_LR_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-LR",
	DateOrder: "DMY",
})

var en_LS_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-LS",
	DateOrder: "DMY",
})

var en_MG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MG",
	DateOrder: "DMY",
})

var en_MH_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MH",
	DateOrder: "MDY",
})

var en_MO_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MO",
	DateOrder: "DMY",
})

var en_MP_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MP",
	DateOrder: "MDY",
})

var en_MS_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MS",
	DateOrder: "DMY",
})

var en_MT_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MT",
	DateOrder: "DMY",
})

var en_MU_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MU",
	DateOrder: "DMY",
})

var en_MW_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MW",
	DateOrder: "DMY",
})

var en_MY_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-MY",
	DateOrder: "DMY",
})

var en_NA_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-NA",
	DateOrder: "DMY",
})

var en_NF_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-NF",
	DateOrder: "DMY",
})

var en_NG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-NG",
	DateOrder: "DMY",
})

var en_NL_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-NL",
	DateOrder: "DMY",
})

var en_NR_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-NR",
	DateOrder: "DMY",
})

var en_NU_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-NU",
	DateOrder: "DMY",
})

var en_NZ_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-NZ",
	DateOrder: "DMY",
})

var en_PG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-PG",
	DateOrder: "DMY",
})

var en_PH_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-PH",
	DateOrder: "DMY",
})

var en_PK_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-PK",
	DateOrder: "DMY",
})

var en_PN_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-PN",
	DateOrder: "DMY",
})

var en_PR_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-PR",
	DateOrder: "MDY",
})

var en_PW_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-PW",
	DateOrder: "DMY",
})

var en_RW_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-RW",
	DateOrder: "DMY",
})

var en_SB_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SB",
	DateOrder: "DMY",
})

var en_SC_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SC",
	DateOrder: "DMY",
})

var en_SD_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SD",
	DateOrder: "DMY",
})

var en_SE_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SE",
	DateOrder: "YMD",
})

var en_SG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SG",
	DateOrder: "DMY",
	Translations: map[string]string{
		"this mth": "0 month ago",
		"last mth": "1 month ago",
		"next mth": "in 1 month",
		"gmt":      "gmt",
		"mth":      "month",
		"utc":      "utc",
		"am":       "am",
		"pm":       "pm",
		" ":        " ",
		"+":        "+",
		"-":        "-",
		".":        ".",
		"/":        "/",
		":":        ":",
		"z":        "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) mth ago`), "$1 month ago"},
		{regexp.MustCompile(`(?i)in (\d+) mth`), "in $1 month"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ mth ago|in \d+ mth)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ mth ago|in \d+ mth)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(last mth|next mth|this mth|gmt|mth|utc|\+|\.|am|pm| |-|/|:|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var en_SH_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SH",
	DateOrder: "DMY",
})

var en_SI_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SI",
	DateOrder: "DMY",
})

var en_SL_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SL",
	DateOrder: "DMY",
})

var en_SS_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SS",
	DateOrder: "DMY",
})

var en_SX_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SX",
	DateOrder: "DMY",
})

var en_SZ_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-SZ",
	DateOrder: "DMY",
})

var en_TC_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-TC",
	DateOrder: "DMY",
})

var en_TK_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-TK",
	DateOrder: "DMY",
})

var en_TO_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-TO",
	DateOrder: "DMY",
})

var en_TT_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-TT",
	DateOrder: "DMY",
})

var en_TV_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-TV",
	DateOrder: "DMY",
})

var en_TZ_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-TZ",
	DateOrder: "DMY",
})

var en_UG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-UG",
	DateOrder: "DMY",
})

var en_UM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-UM",
	DateOrder: "MDY",
})

var en_VC_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-VC",
	DateOrder: "DMY",
})

var en_VG_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-VG",
	DateOrder: "DMY",
})

var en_VI_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-VI",
	DateOrder: "MDY",
})

var en_VU_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-VU",
	DateOrder: "DMY",
})

var en_WS_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-WS",
	DateOrder: "DMY",
})

var en_ZA_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-ZA",
	DateOrder: "YMD",
})

var en_ZM_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-ZM",
	DateOrder: "DMY",
})

var en_ZW_Locale = merge(&en_Locale, LocaleData{
	Name:      "en-ZW",
	DateOrder: "DMY",
})

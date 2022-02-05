// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var yi_Locale = merge(nil, LocaleData{
	Name:      "yi",
	DateOrder: "DMY",
	Charset:   []rune(`ceghiklnorstuwzַָּֿאבגדהוזחטיךכלםמןנסעפץצקרשתײ`),
	Translations: map[string]string{
		"דאנערשטיק": "thursday",
		"נאוועמבער": "november",
		"סעפטעמבער": "september",
		"דעצעמבער":  "december",
		"נאכמיטאג":  "pm",
		"פארמיטאג":  "am",
		"אויגוסט":   "august",
		"אקטאבער":   "october",
		"דינסטיק":   "tuesday",
		"מיטוואך":   "wednesday",
		"סעקונדע":   "second",
		"פעברואר":   "february",
		"זונטיק":    "sunday",
		"יאנואר":    "january",
		"מאנטיק":    "monday",
		"פרײטיק":    "friday",
		"אפריל":     "april",
		"מאנאט":     "month",
		"מינוט":     "minute",
		"אויג":      "august",
		"וואך":      "week",
		"יולי":      "july",
		"יוני":      "june",
		"מערץ":      "march",
		"נאוו":      "november",
		"gmt":       "gmt",
		"utc":       "utc",
		"אפר":       "april",
		"אקט":       "october",
		"דעצ":       "december",
		"טאג":       "day",
		"יאנ":       "january",
		"יאר":       "year",
		"מיי":       "may",
		"סעפ":       "september",
		"פעב":       "february",
		"שבת":       "saturday",
		"שעה":       "hour",
		"am":        "am",
		"pm":        "pm",
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
		"z":         "z",
		"|":         "",
	},
	RelativeType: map[string]string{
		"פארגאנגענעם חודש": "1 month ago",
		"קומענדיקן חודש":   "in 1 month",
		"איבער אכט טאג":    "in 1 week",
		"this minute":      "0 minute ago",
		"איבער א יאר":      "in 1 year",
		"last week":        "1 week ago",
		"this hour":        "0 hour ago",
		"this week":        "0 week ago",
		"דעם חודש":         "0 month ago",
		"פאראיאר":          "1 year ago",
		"הײ יאר":           "0 year ago",
		"היינט":            "0 day ago",
		"מארגן":            "in 1 day",
		"נעכטן":            "1 day ago",
		"now":              "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)אין (\d+) טאג ארום`), "in $1 day"},
		{regexp.MustCompile(`(?i)אין (\d+) טעג ארום`), "in $1 day"},
		{regexp.MustCompile(`(?i)איבער (\d+) חדשים`), "in $1 month"},
		{regexp.MustCompile(`(?i)איבער (\d+) חודש`), "in $1 month"},
		{regexp.MustCompile(`(?i)איבער (\d+) יאר`), "in $1 year"},
		{regexp.MustCompile(`(?i)פאר (\d+) חדשים`), "$1 month ago"},
		{regexp.MustCompile(`(?i)פאר (\d+) חודש`), "$1 month ago"},
		{regexp.MustCompile(`(?i)פאר (\d+) יאר`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(אין \d+ טאג ארום|אין \d+ טעג ארום|איבער \d+ חדשים|איבער \d+ חודש|איבער \d+ יאר|פאר \d+ חדשים|פאר \d+ חודש|פאר \d+ יאר)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(אין \d+ טאג ארום|אין \d+ טעג ארום|איבער \d+ חדשים|איבער \d+ חודש|איבער \d+ יאר|פאר \d+ חדשים|פאר \d+ חודש|פאר \d+ יאר)$`),
	KnownWords:      []string{"פארגאנגענעם חודש", "קומענדיקן חודש", "איבער אכט טאג", "this minute", "איבער א יאר", "last week", "this hour", "this week", "דאנערשטיק", "נאוועמבער", "סעפטעמבער", "דעם חודש", "דעצעמבער", "נאכמיטאג", "פארמיטאג", "אויגוסט", "אקטאבער", "דינסטיק", "מיטוואך", "סעקונדע", "פאראיאר", "פעברואר", "הײ יאר", "זונטיק", "יאנואר", "מאנטיק", "פרײטיק", "אפריל", "היינט", "מאנאט", "מארגן", "מינוט", "נעכטן", "אויג", "וואך", "יולי", "יוני", "מערץ", "נאוו", "gmt", "now", "utc", "אפר", "אקט", "דעצ", "טאג", "יאנ", "יאר", "מיי", "סעפ", "פעב", "שבת", "שעה", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})

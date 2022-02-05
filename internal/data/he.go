// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var he_Locale = merge(nil, LocaleData{
	Name:                  "he",
	DateOrder:             "DMY",
	Charset:               []rune(`-cgtuzאבגדהוזחטיךכלםמןנסעפץצקרשת׳״`),
	SentenceSplitterGroup: 1,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ויום(\z|[^\pL\pM\d]|_)`), "${1}1 יום${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ושבוע(\z|[^\pL\pM\d]|_)`), "${1}1 שבוע${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)וחודש(\z|[^\pL\pM\d]|_)`), "${1}1 חודש${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ושנה(\z|[^\pL\pM\d]|_)`), "${1}1 שנה${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ו(\w+)(\z|[^\pL\pM\d]|_)`), "${1}${2}${3}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)שנתיי?ם(\z|[^\pL\pM\d]|_)`), "${1}2 שנים${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)חודשיים(\z|[^\pL\pM\d]|_)`), "${1}2 חודשים${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)שבועיי?ם(\z|[^\pL\pM\d]|_)`), "${1}2 שבועות${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)יומיי?ם(\z|[^\pL\pM\d]|_)`), "${1}2 ימים${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)שעתיי?ם(\z|[^\pL\pM\d]|_)`), "${1}2 שעות${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)שלשום(\z|[^\pL\pM\d]|_)`), "${1}2 ימים${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)מחר(\z|[^\pL\pM\d]|_)`), "${1}בעוד 1 יום${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)אחר חצות(\z|[^\pL\pM\d]|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)חצות(\z|[^\pL\pM\d]|_)`), "${1}12 am${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)לפנות בוקר(\z|[^\pL\pM\d]|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)בבוקר(\z|[^\pL\pM\d]|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)בצהריי?ם(\z|[^\pL\pM\d]|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)אחרי ה?צהריי?ם(\z|[^\pL\pM\d]|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)אחה"צ(\z|[^\pL\pM\d]|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)לפנות ערב(\z|[^\pL\pM\d]|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)בערב(\z|[^\pL\pM\d]|_)`), "${1}pm${2}"},
	},
	Translations: map[string]string{
		"יום חמישי": "thursday",
		"יום ראשון": "sunday",
		"יום רביעי": "wednesday",
		"יום שלישי": "tuesday",
		"באוקטובר":  "october",
		"יום שישי":  "friday",
		"לאוקטובר":  "october",
		"אוקטובר":   "october",
		"באוגוסט":   "august",
		"בנובמבר":   "november",
		"בסביבות":   "",
		"בספטמבר":   "september",
		"בפברואר":   "february",
		"יום שבת":   "saturday",
		"יום שני":   "monday",
		"לאוגוסט":   "august",
		"לנובמבר":   "november",
		"לספטמבר":   "september",
		"לפברואר":   "february",
		"אוגוסט":    "august",
		"באפריל":    "april",
		"בדצמבר":    "december",
		"בינואר":    "january",
		"בקירוב":    "",
		"חודשים":    "month",
		"יום א׳":    "sunday",
		"יום ב׳":    "monday",
		"יום ג׳":    "tuesday",
		"יום ד׳":    "wednesday",
		"יום ה׳":    "thursday",
		"יום ו׳":    "friday",
		"לאפריל":    "april",
		"לדצמבר":    "december",
		"לינואר":    "january",
		"לפנה״צ":    "am",
		"נובמבר":    "november",
		"ספטמבר":    "september",
		"פברואר":    "february",
		"שבועות":    "week",
		"אחה״צ":     "pm",
		"אפריל":     "april",
		"בחודש":     "month",
		"ביולי":     "july",
		"ביוני":     "june",
		"במארס":     "march",
		"דצמבר":     "december",
		"חמישי":     "thursday",
		"יום א":     "sunday",
		"יום ב":     "monday",
		"יום ג":     "tuesday",
		"יום ד":     "wednesday",
		"יום ה":     "thursday",
		"יום ו":     "friday",
		"ינואר":     "january",
		"ליולי":     "july",
		"ליוני":     "june",
		"למארס":     "march",
		"ראשון":     "sunday",
		"רביעי":     "wednesday",
		"שלישי":     "tuesday",
		"שניות":     "second",
		"שנייה":     "second",
		"אוג׳":      "august",
		"אוק׳":      "october",
		"אפר׳":      "april",
		"ביום":      "day",
		"במאי":      "may",
		"במרס":      "march",
		"במרץ":      "march",
		"בעוד":      "in",
		"בערך":      "",
		"בשנה":      "year",
		"בשעה":      "",
		"דצמ׳":      "december",
		"דקות":      "minute",
		"חודש":      "month",
		"יולי":      "july",
		"יוני":      "june",
		"ימים":      "day",
		"ינו׳":      "january",
		"למאי":      "may",
		"למרס":      "march",
		"למרץ":      "march",
		"לפני":      "ago",
		"מארס":      "march",
		"נוב׳":      "november",
		"ספט׳":      "september",
		"פבר׳":      "february",
		"שבוע":      "week",
		"שישי":      "friday",
		"שניה":      "second",
		"שנים":      "year",
		"שעות":      "hour",
		"gmt":       "gmt",
		"utc":       "utc",
		"אוג":       "august",
		"אוק":       "october",
		"אפר":       "april",
		"דצמ":       "december",
		"דקה":       "minute",
		"דק׳":       "minute",
		"חו׳":       "month",
		"יול":       "july",
		"יום":       "day",
		"יונ":       "june",
		"ינו":       "january",
		"מאי":       "may",
		"מרס":       "march",
		"מרץ":       "march",
		"נוב":       "november",
		"ספט":       "september",
		"עוד":       "in",
		"פבר":       "february",
		"שבת":       "saturday",
		"שב׳":       "week",
		"שנה":       "year",
		"שני":       "monday",
		"שנ׳":       "second",
		"שעה":       "hour",
		"שע׳":       "hour",
		"am":        "am",
		"pm":        "pm",
		"ב-":        "",
		"ה-":        "",
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
		"החודש שעבר": "1 month ago",
		"השבוע שעבר": "1 week ago",
		"השנה שעברה": "1 year ago",
		"החודש הבא":  "in 1 month",
		"השבוע הבא":  "in 1 week",
		"השנה הבאה":  "in 1 year",
		"בדקה זו":    "0 minute ago",
		"בשעה זו":    "0 hour ago",
		"אתמול":      "1 day ago",
		"החודש":      "0 month ago",
		"השבוע":      "0 week ago",
		"עכשיו":      "0 second ago",
		"היום":       "0 day ago",
		"השנה":       "0 year ago",
		"מחר":        "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)בעוד (\d+) חודשים`), "in $1 month"},
		{regexp.MustCompile(`(?i)בעוד (\d+) שבועות`), "in $1 week"},
		{regexp.MustCompile(`(?i)לפני (\d+) חודשים`), "$1 month ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) שבועות`), "$1 week ago"},
		{regexp.MustCompile(`(?i)בעוד (\d+) שניות`), "in $1 second"},
		{regexp.MustCompile(`(?i)לפני (\d+) שניות`), "$1 second ago"},
		{regexp.MustCompile(`(?i)בעוד (\d+) דקות`), "in $1 minute"},
		{regexp.MustCompile(`(?i)בעוד (\d+) ימים`), "in $1 day"},
		{regexp.MustCompile(`(?i)בעוד (\d+) שנים`), "in $1 year"},
		{regexp.MustCompile(`(?i)בעוד (\d+) שעות`), "in $1 hour"},
		{regexp.MustCompile(`(?i)לפני (\d+) דקות`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) ימים`), "$1 day ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) שנים`), "$1 year ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) שעות`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)בעוד (\d+) דק׳`), "in $1 minute"},
		{regexp.MustCompile(`(?i)בעוד (\d+) חו׳`), "in $1 month"},
		{regexp.MustCompile(`(?i)בעוד (\d+) ימ׳`), "in $1 day"},
		{regexp.MustCompile(`(?i)בעוד (\d+) שב׳`), "in $1 week"},
		{regexp.MustCompile(`(?i)בעוד (\d+) שנ׳`), "in $1 second"},
		{regexp.MustCompile(`(?i)בעוד (\d+) שע׳`), "in $1 hour"},
		{regexp.MustCompile(`(?i)בעוד יום (\d+)`), "in $1 day"},
		{regexp.MustCompile(`(?i)לפני (\d+) דק׳`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) חו׳`), "$1 month ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) ימ׳`), "$1 day ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) שב׳`), "$1 week ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) שנ׳`), "$1 second ago"},
		{regexp.MustCompile(`(?i)לפני (\d+) שע׳`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)לפני יום (\d+)`), "$1 day ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(בעוד \d+ חודשים|בעוד \d+ שבועות|לפני \d+ חודשים|לפני \d+ שבועות|בעוד \d+ שניות|לפני \d+ שניות|בעוד \d+ דקות|בעוד \d+ ימים|בעוד \d+ שנים|בעוד \d+ שעות|לפני \d+ דקות|לפני \d+ ימים|לפני \d+ שנים|לפני \d+ שעות|בעוד \d+ דק׳|בעוד \d+ חו׳|בעוד \d+ ימ׳|בעוד \d+ שב׳|בעוד \d+ שנ׳|בעוד \d+ שע׳|בעוד יום \d+|לפני \d+ דק׳|לפני \d+ חו׳|לפני \d+ ימ׳|לפני \d+ שב׳|לפני \d+ שנ׳|לפני \d+ שע׳|לפני יום \d+)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(בעוד \d+ חודשים|בעוד \d+ שבועות|לפני \d+ חודשים|לפני \d+ שבועות|בעוד \d+ שניות|לפני \d+ שניות|בעוד \d+ דקות|בעוד \d+ ימים|בעוד \d+ שנים|בעוד \d+ שעות|לפני \d+ דקות|לפני \d+ ימים|לפני \d+ שנים|לפני \d+ שעות|בעוד \d+ דק׳|בעוד \d+ חו׳|בעוד \d+ ימ׳|בעוד \d+ שב׳|בעוד \d+ שנ׳|בעוד \d+ שע׳|בעוד יום \d+|לפני \d+ דק׳|לפני \d+ חו׳|לפני \d+ ימ׳|לפני \d+ שב׳|לפני \d+ שנ׳|לפני \d+ שע׳|לפני יום \d+)$`),
	KnownWords:      []string{"החודש שעבר", "השבוע שעבר", "השנה שעברה", "החודש הבא", "השבוע הבא", "השנה הבאה", "יום חמישי", "יום ראשון", "יום רביעי", "יום שלישי", "באוקטובר", "יום שישי", "לאוקטובר", "אוקטובר", "באוגוסט", "בדקה זו", "בנובמבר", "בסביבות", "בספטמבר", "בפברואר", "בשעה זו", "יום שבת", "יום שני", "לאוגוסט", "לנובמבר", "לספטמבר", "לפברואר", "אוגוסט", "באפריל", "בדצמבר", "בינואר", "בקירוב", "חודשים", "יום א׳", "יום ב׳", "יום ג׳", "יום ד׳", "יום ה׳", "יום ו׳", "לאפריל", "לדצמבר", "לינואר", "לפנה״צ", "נובמבר", "ספטמבר", "פברואר", "שבועות", "אחה״צ", "אפריל", "אתמול", "בחודש", "ביולי", "ביוני", "במארס", "דצמבר", "החודש", "השבוע", "חמישי", "יום א", "יום ב", "יום ג", "יום ד", "יום ה", "יום ו", "ינואר", "ליולי", "ליוני", "למארס", "עכשיו", "ראשון", "רביעי", "שלישי", "שניות", "שנייה", "אוג׳", "אוק׳", "אפר׳", "ביום", "במאי", "במרס", "במרץ", "בעוד", "בערך", "בשנה", "בשעה", "דצמ׳", "דקות", "היום", "השנה", "חודש", "יולי", "יוני", "ימים", "ינו׳", "למאי", "למרס", "למרץ", "לפני", "מארס", "נוב׳", "ספט׳", "פבר׳", "שבוע", "שישי", "שניה", "שנים", "שעות", "gmt", "utc", "אוג", "אוק", "אפר", "דצמ", "דקה", "דק׳", "חו׳", "יול", "יום", "יונ", "ינו", "מאי", "מחר", "מרס", "מרץ", "נוב", "ספט", "עוד", "פבר", "שבת", "שב׳", "שנה", "שני", "שנ׳", "שעה", "שע׳", "am", "pm", "ב-", "ה-", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})

// Code is generated by script; DO NOT EDIT.

package data

import "github.com/markusmobius/go-dateparser/internal/regexp"

var (
	nn_Locale LocaleData
)

func init() {
	nn_Locale = merge(nil, LocaleData{
		Name:      "nn",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuvwxyzåø`),
		Translations: map[string][]string{
			"ettermiddag": {"pm"},
			"formiddag":   {"am"},
			"september":   {"september"},
			"desember":    {"december"},
			"november":    {"november"},
			"februar":     {"february"},
			"laurdag":     {"saturday"},
			"oktober":     {"october"},
			"torsdag":     {"thursday"},
			"august":      {"august"},
			"fredag":      {"friday"},
			"januar":      {"january"},
			"mandag":      {"monday"},
			"minutt":      {"minute"},
			"onsdag":      {"wednesday"},
			"sekund":      {"second"},
			"søndag":      {"sunday"},
			"tysdag":      {"tuesday"},
			"april":       {"april"},
			"manad":       {"month"},
			"juli":        {"july"},
			"juni":        {"june"},
			"mars":        {"march"},
			"time":        {"hour"},
			"veke":        {"week"},
			"apr":         {"april"},
			"aug":         {"august"},
			"dag":         {"day"},
			"des":         {"december"},
			"feb":         {"february"},
			"fre":         {"friday"},
			"gmt":         {"gmt"},
			"jan":         {"january"},
			"jul":         {"july"},
			"jun":         {"june"},
			"lau":         {"saturday"},
			"mai":         {"may"},
			"man":         {"monday"},
			"mar":         {"march"},
			"nov":         {"november"},
			"okt":         {"october"},
			"ons":         {"wednesday"},
			"sep":         {"september"},
			"søn":         {"sunday"},
			"tor":         {"thursday"},
			"tys":         {"tuesday"},
			"utc":         {"utc"},
			"am":          {"am"},
			"ar":          {"year"},
			"em":          {"pm"},
			"fm":          {"am"},
			"fr":          {"friday"},
			"la":          {"saturday"},
			"ma":          {"monday"},
			"on":          {"wednesday"},
			"pm":          {"pm"},
			"sø":          {"sunday"},
			"to":          {"thursday"},
			"ty":          {"tuesday"},
			" ":           {" "},
			"'":           {""},
			"+":           {"+"},
			",":           {""},
			"-":           {"-"},
			".":           {"."},
			"/":           {"/"},
			":":           {":"},
			";":           {""},
			"@":           {""},
			"[":           {""},
			"]":           {""},
			"z":           {"z"},
			"|":           {""},
		},
		RelativeType: map[string]string{
			"this minute": "0 minute ago",
			"last month":  "1 month ago",
			"next month":  "in 1 month",
			"this month":  "0 month ago",
			"last week":   "1 week ago",
			"last year":   "1 year ago",
			"next week":   "in 1 week",
			"next year":   "in 1 year",
			"this hour":   "0 hour ago",
			"this week":   "0 week ago",
			"this year":   "0 year ago",
			"i morgon":    "in 1 day",
			"i dag":       "0 day ago",
			"i gar":       "1 day ago",
			"now":         "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) minutter siden`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sekunder siden`), "$1 second ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) maneder siden`), "$1 month ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) minutt siden`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sekund siden`), "$1 second ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) maned siden`), "$1 month ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) timer siden`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) døgn siden`), "$1 day ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) time siden`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) uker siden`), "$1 week ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) uke siden`), "$1 week ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) ar siden`), "$1 year ago"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) minutter`), "in $1 minute"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sekunder`), "in $1 second"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) maneder`), "in $1 month"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) minutt`), "in $1 minute"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sekund`), "in $1 second"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) maned`), "in $1 month"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) timer`), "in $1 hour"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) døgn`), "in $1 day"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) time`), "in $1 hour"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) uker`), "in $1 week"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) uke`), "in $1 week"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) ar`), "in $1 year"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(for \d+[.,]?\d* minutter siden|for \d+[.,]?\d* sekunder siden|for \d+[.,]?\d* maneder siden|for \d+[.,]?\d* minutt siden|for \d+[.,]?\d* sekund siden|for \d+[.,]?\d* maned siden|for \d+[.,]?\d* timer siden|for \d+[.,]?\d* døgn siden|for \d+[.,]?\d* time siden|for \d+[.,]?\d* uker siden|for \d+[.,]?\d* uke siden|for \d+[.,]?\d* ar siden|om \d+[.,]?\d* minutter|om \d+[.,]?\d* sekunder|om \d+[.,]?\d* maneder|om \d+[.,]?\d* minutt|om \d+[.,]?\d* sekund|om \d+[.,]?\d* maned|om \d+[.,]?\d* timer|om \d+[.,]?\d* døgn|om \d+[.,]?\d* time|om \d+[.,]?\d* uker|om \d+[.,]?\d* uke|om \d+[.,]?\d* ar)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(for \d+[.,]?\d* minutter siden|for \d+[.,]?\d* sekunder siden|for \d+[.,]?\d* maneder siden|for \d+[.,]?\d* minutt siden|for \d+[.,]?\d* sekund siden|for \d+[.,]?\d* maned siden|for \d+[.,]?\d* timer siden|for \d+[.,]?\d* døgn siden|for \d+[.,]?\d* time siden|for \d+[.,]?\d* uker siden|for \d+[.,]?\d* uke siden|for \d+[.,]?\d* ar siden|om \d+[.,]?\d* minutter|om \d+[.,]?\d* sekunder|om \d+[.,]?\d* maneder|om \d+[.,]?\d* minutt|om \d+[.,]?\d* sekund|om \d+[.,]?\d* maned|om \d+[.,]?\d* timer|om \d+[.,]?\d* døgn|om \d+[.,]?\d* time|om \d+[.,]?\d* uker|om \d+[.,]?\d* uke|om \d+[.,]?\d* ar)$`),
		KnownWords:      []string{"ettermiddag", "this minute", "last month", "next month", "this month", "formiddag", "last week", "last year", "next week", "next year", "september", "this hour", "this week", "this year", "desember", "i morgon", "november", "februar", "laurdag", "oktober", "torsdag", "august", "fredag", "januar", "mandag", "minutt", "onsdag", "sekund", "søndag", "tysdag", "april", "i dag", "i gar", "manad", "juli", "juni", "mars", "time", "veke", "apr", "aug", "dag", "des", "feb", "fre", "gmt", "jan", "jul", "jun", "lau", "mai", "man", "mar", "nov", "now", "okt", "ons", "sep", "søn", "tor", "tys", "utc", "am", "ar", "em", "fm", "fr", "la", "ma", "on", "pm", "sø", "to", "ty", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}

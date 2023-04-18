// Code is generated by script; DO NOT EDIT.

package data

import "github.com/markusmobius/go-dateparser/internal/regexp"

var (
	be_Locale LocaleData
)

func init() {
	be_Locale = merge(nil, LocaleData{
		Name:                  "be",
		DateOrder:             "DMY",
		Charset:               []rune(`cgtuzабвгдежзийклмнопрстухцчыьэюяёіў`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^гадзіна(\z|[^\pL\pM\d]|_)`), "${1}1 гадзіна${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)гадзіну(\z|[^\pL\pM\d]|_)`), "${1}1 гадзіну${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^хвіліну(\z|[^\pL\pM\d]|_)`), "${1}1 хвіліну${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^секунду(\z|[^\pL\pM\d]|_)`), "${1}1 секунду${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)некалькі секунд(\z|[^\pL\pM\d]|_)`), "${1}44 секунды${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)некалькі хвілін(\z|[^\pL\pM\d]|_)`), "${1}2 хвіліны${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d*)\s*гадзін\s(\d+[.,]?\d*)\s*хвілін(\z|[^\pL\pM\d]|_)`), "${1}${2}:${3}${4}"},
		},
		Translations: map[string][]string{
			"кастрычніка": {"october"},
			"кастрычнік":  {"october"},
			"на працягу":  {"in"},
			"панядзелак":  {"monday"},
			"таму назад":  {"ago"},
			"красавіка":   {"april"},
			"лістапада":   {"november"},
			"верасень":    {"september"},
			"красавік":    {"april"},
			"лістапад":    {"november"},
			"сакавіка":    {"march"},
			"студзень":    {"january"},
			"студзеня":    {"january"},
			"ауторак":     {"tuesday"},
			"верасня":     {"september"},
			"гадзіна":     {"hour"},
			"гадзіну":     {"hour"},
			"гадзіны":     {"hour"},
			"жнівень":     {"august"},
			"жнівеня":     {"august"},
			"месяцау":     {"month"},
			"нядзеля":     {"sunday"},
			"пятніца":     {"friday"},
			"сакавік":     {"march"},
			"секунда":     {"second"},
			"секунду":     {"second"},
			"секунды":     {"second"},
			"снежань":     {"december"},
			"тыдзень":     {"week"},
			"хвіліна":     {"minute"},
			"хвіліну":     {"minute"},
			"хвіліны":     {"minute"},
			"чэрвень":     {"june"},
			"чэрвеня":     {"june"},
			"гадзін":      {"hour"},
			"жніуня":      {"august"},
			"лютага":      {"february"},
			"ліпень":      {"july"},
			"ліпеня":      {"july"},
			"месяца":      {"month"},
			"месяцы":      {"month"},
			"секунд":      {"second"},
			"серада":      {"wednesday"},
			"снежня":      {"december"},
			"субота":      {"saturday"},
			"трауня":      {"may"},
			"тыдняу":      {"week"},
			"хвілін":      {"minute"},
			"чацвер":      {"thursday"},
			"гадоу":       {"year"},
			"дзень":       {"day"},
			"месяц":       {"month"},
			"назад":       {"ago"},
			"тыдня":       {"week"},
			"тыдні":       {"week"},
			"гадз":        {"hour"},
			"гады":        {"year"},
			"года":        {"year"},
			"дзен":        {"day"},
			"каля":        {""},
			"люты":        {"february"},
			"таму":        {"ago"},
			"хвіл":        {"minute"},
			"gmt":         {"gmt"},
			"utc":         {"utc"},
			"аут":         {"tuesday"},
			"вер":         {"september"},
			"врс":         {"september"},
			"год":         {"year"},
			"дні":         {"day"},
			"жнв":         {"august"},
			"жні":         {"august"},
			"кас":         {"october"},
			"кра":         {"april"},
			"крс":         {"april"},
			"кст":         {"october"},
			"лют":         {"february"},
			"ліп":         {"july"},
			"ліс":         {"november"},
			"маи":         {"may"},
			"мая":         {"may"},
			"мес":         {"month"},
			"няд":         {"sunday"},
			"пнд":         {"monday"},
			"пят":         {"friday"},
			"сак":         {"march"},
			"сек":         {"second"},
			"сне":         {"december"},
			"снж":         {"december"},
			"стд":         {"january"},
			"сту":         {"january"},
			"суб":         {"saturday"},
			"тра":         {"may"},
			"тыд":         {"week"},
			"чцв":         {"thursday"},
			"чэр":         {"june"},
			"am":          {"am"},
			"pm":          {"pm"},
			"ау":          {"tuesday"},
			"нд":          {"sunday"},
			"пн":          {"monday"},
			"пт":          {"friday"},
			"сб":          {"saturday"},
			"ср":          {"wednesday"},
			"хв":          {"minute"},
			"чв":          {"thursday"},
			"чц":          {"thursday"},
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
			"г":           {"year"},
			"д":           {"day"},
			"с":           {"second"},
			"у":           {""},
			"і":           {""},
		},
		RelativeType: map[string]string{
			"на наступным тыдні": "in 1 week",
			"у наступным месяцы": "in 1 month",
			"у наступным годзе":  "in 1 year",
			"на мінулым тыдні":   "1 week ago",
			"у мінулым месяцы":   "1 month ago",
			"у мінулым годзе":    "1 year ago",
			"на гэтым тыдні":     "0 week ago",
			"у гэту гадзіну":     "0 hour ago",
			"у гэту хвіліну":     "0 minute ago",
			"у гэтым месяцы":     "0 month ago",
			"у гэтым годзе":      "0 year ago",
			"пазаучора":          "2 day ago",
			"заутра":             "in 1 day",
			"сення":              "0 day ago",
			"учора":              "1 day ago",
			"цяпер":              "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) гадзіну таму`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) гадзіны таму`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) секунду таму`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) секунды таму`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) тыдзень таму`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) хвіліну таму`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) хвіліны таму`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) гадзіну`), "in $1 hour"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) гадзіны`), "in $1 hour"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) секунду`), "in $1 second"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) секунды`), "in $1 second"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) тыдзень`), "in $1 week"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) хвіліну`), "in $1 minute"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) хвіліны`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) месяца таму`), "$1 month ago"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) месяца`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) дзень таму`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) месяц таму`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) тыдня таму`), "$1 week ago"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) дзень`), "in $1 day"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) месяц`), "in $1 month"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) тыдня`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) гадз таму`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) года таму`), "$1 year ago"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) гадз`), "in $1 hour"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) года`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) год таму`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) дня таму`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) мес таму`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) тыд таму`), "$1 week ago"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) год`), "in $1 year"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) дня`), "in $1 day"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) мес`), "in $1 month"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) тыд`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) хв таму`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) хв`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) г таму`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) д таму`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) с таму`), "$1 second ago"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) г`), "in $1 year"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) д`), "in $1 day"},
			{regexp.MustCompile(`(?i)праз (\d+[.,]?\d*) с`), "in $1 second"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* гадзіну таму|\d+[.,]?\d* гадзіны таму|\d+[.,]?\d* секунду таму|\d+[.,]?\d* секунды таму|\d+[.,]?\d* тыдзень таму|\d+[.,]?\d* хвіліну таму|\d+[.,]?\d* хвіліны таму|праз \d+[.,]?\d* гадзіну|праз \d+[.,]?\d* гадзіны|праз \d+[.,]?\d* секунду|праз \d+[.,]?\d* секунды|праз \d+[.,]?\d* тыдзень|праз \d+[.,]?\d* хвіліну|праз \d+[.,]?\d* хвіліны|\d+[.,]?\d* месяца таму|праз \d+[.,]?\d* месяца|\d+[.,]?\d* дзень таму|\d+[.,]?\d* месяц таму|\d+[.,]?\d* тыдня таму|праз \d+[.,]?\d* дзень|праз \d+[.,]?\d* месяц|праз \d+[.,]?\d* тыдня|\d+[.,]?\d* гадз таму|\d+[.,]?\d* года таму|праз \d+[.,]?\d* гадз|праз \d+[.,]?\d* года|\d+[.,]?\d* год таму|\d+[.,]?\d* дня таму|\d+[.,]?\d* мес таму|\d+[.,]?\d* тыд таму|праз \d+[.,]?\d* год|праз \d+[.,]?\d* дня|праз \d+[.,]?\d* мес|праз \d+[.,]?\d* тыд|\d+[.,]?\d* хв таму|праз \d+[.,]?\d* хв|\d+[.,]?\d* г таму|\d+[.,]?\d* д таму|\d+[.,]?\d* с таму|праз \d+[.,]?\d* г|праз \d+[.,]?\d* д|праз \d+[.,]?\d* с)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* гадзіну таму|\d+[.,]?\d* гадзіны таму|\d+[.,]?\d* секунду таму|\d+[.,]?\d* секунды таму|\d+[.,]?\d* тыдзень таму|\d+[.,]?\d* хвіліну таму|\d+[.,]?\d* хвіліны таму|праз \d+[.,]?\d* гадзіну|праз \d+[.,]?\d* гадзіны|праз \d+[.,]?\d* секунду|праз \d+[.,]?\d* секунды|праз \d+[.,]?\d* тыдзень|праз \d+[.,]?\d* хвіліну|праз \d+[.,]?\d* хвіліны|\d+[.,]?\d* месяца таму|праз \d+[.,]?\d* месяца|\d+[.,]?\d* дзень таму|\d+[.,]?\d* месяц таму|\d+[.,]?\d* тыдня таму|праз \d+[.,]?\d* дзень|праз \d+[.,]?\d* месяц|праз \d+[.,]?\d* тыдня|\d+[.,]?\d* гадз таму|\d+[.,]?\d* года таму|праз \d+[.,]?\d* гадз|праз \d+[.,]?\d* года|\d+[.,]?\d* год таму|\d+[.,]?\d* дня таму|\d+[.,]?\d* мес таму|\d+[.,]?\d* тыд таму|праз \d+[.,]?\d* год|праз \d+[.,]?\d* дня|праз \d+[.,]?\d* мес|праз \d+[.,]?\d* тыд|\d+[.,]?\d* хв таму|праз \d+[.,]?\d* хв|\d+[.,]?\d* г таму|\d+[.,]?\d* д таму|\d+[.,]?\d* с таму|праз \d+[.,]?\d* г|праз \d+[.,]?\d* д|праз \d+[.,]?\d* с)$`),
		KnownWords:      []string{"на наступным тыдні", "у наступным месяцы", "у наступным годзе", "на мінулым тыдні", "у мінулым месяцы", "у мінулым годзе", "на гэтым тыдні", "у гэту гадзіну", "у гэту хвіліну", "у гэтым месяцы", "у гэтым годзе", "кастрычніка", "кастрычнік", "на працягу", "панядзелак", "таму назад", "красавіка", "лістапада", "пазаучора", "верасень", "красавік", "лістапад", "сакавіка", "студзень", "студзеня", "ауторак", "верасня", "гадзіна", "гадзіну", "гадзіны", "жнівень", "жнівеня", "месяцау", "нядзеля", "пятніца", "сакавік", "секунда", "секунду", "секунды", "снежань", "тыдзень", "хвіліна", "хвіліну", "хвіліны", "чэрвень", "чэрвеня", "гадзін", "жніуня", "заутра", "лютага", "ліпень", "ліпеня", "месяца", "месяцы", "секунд", "серада", "снежня", "субота", "трауня", "тыдняу", "хвілін", "чацвер", "гадоу", "дзень", "месяц", "назад", "сення", "тыдня", "тыдні", "учора", "цяпер", "гадз", "гады", "года", "дзен", "каля", "люты", "таму", "хвіл", "gmt", "utc", "аут", "вер", "врс", "год", "дні", "жнв", "жні", "кас", "кра", "крс", "кст", "лют", "ліп", "ліс", "маи", "мая", "мес", "няд", "пнд", "пят", "сак", "сек", "сне", "снж", "стд", "сту", "суб", "тра", "тыд", "чцв", "чэр", "am", "pm", "ау", "нд", "пн", "пт", "сб", "ср", "хв", "чв", "чц", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "г", "д", "с", "у", "і"},
	})
}

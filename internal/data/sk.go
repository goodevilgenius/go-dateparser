// Code is generated by script; DO NOT EDIT.

package data

import "github.com/markusmobius/go-dateparser/internal/regexp"

var (
	sk_Locale LocaleData
)

func init() {
	sk_Locale = merge(nil, LocaleData{
		Name:                  "sk",
		DateOrder:             "DMY",
		Charset:               []rune(`bcdefghijklnorstuvyzáíóúýčľňšž`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)pred sekundou(\z|[^\pL\pM\d]|_)`), "${1}pred 1 sekundou${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)pred minutou(\z|[^\pL\pM\d]|_)`), "${1}pred 1 minutou${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)pred hodinou(\z|[^\pL\pM\d]|_)`), "${1}pred 1 hodinou${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)pred tyzdnom(\z|[^\pL\pM\d]|_)`), "${1}pred 1 tyzdnom${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)pred mesiacom(\z|[^\pL\pM\d]|_)`), "${1}pred 1 mesiacom${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)pred rokom(\z|[^\pL\pM\d]|_)`), "${1}pred 1 rokom${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)pred pol rokom(\z|[^\pL\pM\d]|_)`), "${1}pred 6 mesiacmi${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)o sekundu(\z|[^\pL\pM\d]|_)`), "${1}o 1 sekundu${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)o minutu(\z|[^\pL\pM\d]|_)`), "${1}o 1 minutu${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)o hodinu(\z|[^\pL\pM\d]|_)`), "${1}o 1 hodinu${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)o tyzden(\z|[^\pL\pM\d]|_)`), "${1}o 1 tyzden${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)o mesiac(\z|[^\pL\pM\d]|_)`), "${1}o 1 mesiac${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)o rok(\z|[^\pL\pM\d]|_)`), "${1}o 1 rok${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)o pol roka(\z|[^\pL\pM\d]|_)`), "${1}o 6 mesiacov${2}"},
		},
		Translations: map[string][]string{
			"september": {"september"},
			"septembra": {"september"},
			"december":  {"december"},
			"decembra":  {"december"},
			"februara":  {"february"},
			"november":  {"november"},
			"novembra":  {"november"},
			"pondelok":  {"monday"},
			"augusta":   {"august"},
			"februar":   {"february"},
			"januara":   {"january"},
			"oktober":   {"october"},
			"oktobra":   {"october"},
			"sekunda":   {"second"},
			"stvrtok":   {"thursday"},
			"aprila":    {"april"},
			"august":    {"august"},
			"hodina":    {"hour"},
			"januar":    {"january"},
			"mesiac":    {"month"},
			"minuta":    {"minute"},
			"nedela":    {"sunday"},
			"nedelu":    {"sunday"},
			"piatok":    {"friday"},
			"sobota":    {"saturday"},
			"sobotu":    {"saturday"},
			"streda":    {"wednesday"},
			"stredu":    {"wednesday"},
			"tyzden":    {"week"},
			"utorok":    {"tuesday"},
			"april":     {"april"},
			"marca":     {"march"},
			"marec":     {"march"},
			"jula":      {"july"},
			"juna":      {"june"},
			"maja":      {"may"},
			"pred":      {"ago"},
			"apr":       {"april"},
			"aug":       {"august"},
			"dec":       {"december"},
			"den":       {"day"},
			"feb":       {"february"},
			"gmt":       {"gmt"},
			"jan":       {"january"},
			"jul":       {"july"},
			"jun":       {"june"},
			"maj":       {"may"},
			"mar":       {"march"},
			"mes":       {"month"},
			"min":       {"minute"},
			"nov":       {"november"},
			"okt":       {"october"},
			"rok":       {"year"},
			"sep":       {"september"},
			"tyz":       {"week"},
			"utc":       {"utc"},
			"am":        {"am"},
			"ne":        {"sunday"},
			"pi":        {"friday"},
			"pm":        {"pm"},
			"po":        {"monday"},
			"so":        {"saturday"},
			"st":        {"thursday", "wednesday"},
			"ut":        {"tuesday"},
			"vo":        {""},
			"za":        {"in"},
			" ":         {" "},
			"'":         {""},
			"+":         {"+"},
			",":         {""},
			"-":         {"-"},
			".":         {"."},
			"/":         {"/"},
			":":         {":"},
			";":         {""},
			"@":         {""},
			"[":         {""},
			"]":         {""},
			"d":         {"day"},
			"h":         {"hour"},
			"o":         {"in"},
			"r":         {"year"},
			"s":         {"second"},
			"v":         {""},
			"z":         {"z"},
			"|":         {""},
		},
		RelativeType: map[string]string{
			"v tejto hodine": "0 hour ago",
			"v tejto minute": "0 minute ago",
			"buduci mesiac":  "in 1 month",
			"buduci tyzden":  "in 1 week",
			"minuly mesiac":  "1 month ago",
			"minuly tyzden":  "1 week ago",
			"tento mesiac":   "0 month ago",
			"tento tyzden":   "0 week ago",
			"buduci rok":     "in 1 year",
			"minuly rok":     "1 year ago",
			"predvcerom":     "2 day ago",
			"tento rok":      "0 year ago",
			"pozajtra":       "in 2 days",
			"zajtra":         "in 1 day",
			"teraz":          "0 second ago",
			"vcera":          "1 day ago",
			"dnes":           "0 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) sekundami`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) hodinami`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) mesiacmi`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) mesiacom`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) minutami`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) sekundou`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) tyzdnami`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) hodinou`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) minutou`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) tyzdnom`), "$1 week ago"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) mesiacov`), "in $1 month"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) dnami`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) rokmi`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) rokom`), "$1 year ago"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) mesiace`), "in ${1} month"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) sekundu`), "in $1 second"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) sekundy`), "in ${1} second"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) tyzdnov`), "in $1 week"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) dnom`), "$1 day ago"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) hodinu`), "in $1 hour"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) hodiny`), "in ${1} hour"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) mesiac`), "in $1 month"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) minutu`), "in $1 minute"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) minuty`), "in ${1} minute"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) sekund`), "in $1 second"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) tyzden`), "in $1 week"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) tyzdne`), "in ${1} week"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) mes`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) tyz`), "$1 week ago"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) hodin`), "in $1 hour"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) minut`), "in $1 minute"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) rokov`), "in $1 year"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) roky`), "in ${1} year"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) d`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) h`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) r`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) s`), "$1 second ago"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) den`), "in $1 day"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) dni`), "in ${1} day"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) mes`), "in $1 month"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) rok`), "in $1 year"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) tyz`), "in $1 week"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) d`), "in $1 day"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) h`), "in $1 hour"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) r`), "in $1 year"},
			{regexp.MustCompile(`(?i)o (\d+[.,]?\d*) s`), "in $1 second"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pred \d+[.,]?\d* sekundami|pred \d+[.,]?\d* hodinami|pred \d+[.,]?\d* mesiacmi|pred \d+[.,]?\d* mesiacom|pred \d+[.,]?\d* minutami|pred \d+[.,]?\d* sekundou|pred \d+[.,]?\d* tyzdnami|pred \d+[.,]?\d* hodinou|pred \d+[.,]?\d* minutou|pred \d+[.,]?\d* tyzdnom|o \d+[.,]?\d* mesiacov|pred \d+[.,]?\d* dnami|pred \d+[.,]?\d* rokmi|pred \d+[.,]?\d* rokom|o \d+[.,]?\d* mesiace|o \d+[.,]?\d* sekundu|o \d+[.,]?\d* sekundy|o \d+[.,]?\d* tyzdnov|pred \d+[.,]?\d* dnom|o \d+[.,]?\d* hodinu|o \d+[.,]?\d* hodiny|o \d+[.,]?\d* mesiac|o \d+[.,]?\d* minutu|o \d+[.,]?\d* minuty|o \d+[.,]?\d* sekund|o \d+[.,]?\d* tyzden|o \d+[.,]?\d* tyzdne|pred \d+[.,]?\d* mes|pred \d+[.,]?\d* min|pred \d+[.,]?\d* tyz|o \d+[.,]?\d* hodin|o \d+[.,]?\d* minut|o \d+[.,]?\d* rokov|o \d+[.,]?\d* roky|pred \d+[.,]?\d* d|pred \d+[.,]?\d* h|pred \d+[.,]?\d* r|pred \d+[.,]?\d* s|o \d+[.,]?\d* den|o \d+[.,]?\d* dni|o \d+[.,]?\d* mes|o \d+[.,]?\d* min|o \d+[.,]?\d* rok|o \d+[.,]?\d* tyz|o \d+[.,]?\d* d|o \d+[.,]?\d* h|o \d+[.,]?\d* r|o \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(pred \d+[.,]?\d* sekundami|pred \d+[.,]?\d* hodinami|pred \d+[.,]?\d* mesiacmi|pred \d+[.,]?\d* mesiacom|pred \d+[.,]?\d* minutami|pred \d+[.,]?\d* sekundou|pred \d+[.,]?\d* tyzdnami|pred \d+[.,]?\d* hodinou|pred \d+[.,]?\d* minutou|pred \d+[.,]?\d* tyzdnom|o \d+[.,]?\d* mesiacov|pred \d+[.,]?\d* dnami|pred \d+[.,]?\d* rokmi|pred \d+[.,]?\d* rokom|o \d+[.,]?\d* mesiace|o \d+[.,]?\d* sekundu|o \d+[.,]?\d* sekundy|o \d+[.,]?\d* tyzdnov|pred \d+[.,]?\d* dnom|o \d+[.,]?\d* hodinu|o \d+[.,]?\d* hodiny|o \d+[.,]?\d* mesiac|o \d+[.,]?\d* minutu|o \d+[.,]?\d* minuty|o \d+[.,]?\d* sekund|o \d+[.,]?\d* tyzden|o \d+[.,]?\d* tyzdne|pred \d+[.,]?\d* mes|pred \d+[.,]?\d* min|pred \d+[.,]?\d* tyz|o \d+[.,]?\d* hodin|o \d+[.,]?\d* minut|o \d+[.,]?\d* rokov|o \d+[.,]?\d* roky|pred \d+[.,]?\d* d|pred \d+[.,]?\d* h|pred \d+[.,]?\d* r|pred \d+[.,]?\d* s|o \d+[.,]?\d* den|o \d+[.,]?\d* dni|o \d+[.,]?\d* mes|o \d+[.,]?\d* min|o \d+[.,]?\d* rok|o \d+[.,]?\d* tyz|o \d+[.,]?\d* d|o \d+[.,]?\d* h|o \d+[.,]?\d* r|o \d+[.,]?\d* s)$`),
		KnownWords:      []string{"v tejto hodine", "v tejto minute", "buduci mesiac", "buduci tyzden", "minuly mesiac", "minuly tyzden", "tento mesiac", "tento tyzden", "buduci rok", "minuly rok", "predvcerom", "september", "septembra", "tento rok", "december", "decembra", "februara", "november", "novembra", "pondelok", "pozajtra", "augusta", "februar", "januara", "oktober", "oktobra", "sekunda", "stvrtok", "aprila", "august", "hodina", "januar", "mesiac", "minuta", "nedela", "nedelu", "piatok", "sobota", "sobotu", "streda", "stredu", "tyzden", "utorok", "zajtra", "april", "marca", "marec", "teraz", "vcera", "dnes", "jula", "juna", "maja", "pred", "apr", "aug", "dec", "den", "feb", "gmt", "jan", "jul", "jun", "maj", "mar", "mes", "min", "nov", "okt", "rok", "sep", "tyz", "utc", "am", "ne", "pi", "pm", "po", "so", "st", "ut", "vo", "za", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "h", "o", "r", "s", "v", "z", "|"},
	})
}
